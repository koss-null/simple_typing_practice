package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
)

// Service settles orders by calling a third party API.
type Service struct {
    repo OrderRepository
    api  ThirdParty
}

// Order represents a funds withdrawal intent.
// Add missing fields if needed.
type Order struct {
    // ID is an internal identifier of the order.
    ID string
    // SettlementID is an identifier of order's settlement in a third party system.
    SettlementID string
    // Amount of funds to withdraw.
    Amount uint
    // Status of the fund withdrawal intent.
    Status Status
}

// Status of an Order.
type Status int

// Add missing statuses if needed.
const (
    StatusPending Status = iota
    StatusSuccess
)

// Storage allows to execute SQL transactions.
// For example, a service might need to call different methods within the same database transaction.
type Storage interface {
    // Transact executes a function where transaction atomicity on the database is guaranteed.
    // If the function is successfully completed, the changes are committed to the database.
    // If there is an error, the changes are rolled back.
    Transact(atomic func(tx *sql.Tx) error) error
}

// OrderRepository is used to store and access orders.
// Add missing methods if needed.
type OrderRepository interface {
    Storage
    // FindByID looks up an order by its ID.
    FindByID(tx *sql.Tx, id string) (Order, error)
    // ListPending fetches all pending orders.
    ListPending() ([]Order, error)
    // MarkSuccess changes specified order status to success.
    MarkSuccess(tx *sql.Tx, order Order) error
}

// ThirdParty is used to settle orders by calling a third party API.
type ThirdParty interface {
    // Settle an order in a third party system and return ID of the settlement.
    Settle(order Order) (string, error)
}

type fakeOrderRepository struct{}

func (r *fakeOrderRepository) Transact(atomic func(tx *sql.Tx) error) error {
    return atomic(&sql.Tx{})
}

func (r *fakeOrderRepository) FindByID(tx *sql.Tx, id string) (Order, error) {
    return Order{}, nil
}

func (r *fakeOrderRepository) ListPending() ([]Order, error) {
    return []Order{}, nil
}

func (r *fakeOrderRepository) MarkSuccess(tx *sql.Tx, order Order) error {
    return nil
}

type fakeThirdParty struct{}

func (a *fakeThirdParty) Settle(order Order) (string, error) {
    return "", nil
}

// processOrders settles pending orders until an error occurs.
// This method is thread-safe. Meaning you can scale processing by running multiple
// instances of this function in parallel.
func (s *Service) processOrders() error {
    orders, err := s.repo.ListPending()
    if err != nil {
        return fmt.Errorf("error fetching pending orders: %w", err)
    }

    for i := range orders {
        err = s.settleOrder(orders[i])
        if err != nil {
            return fmt.Errorf("error processing order %s: %w", orders[i].ID, err)
        }
    }

    return nil
}

// Pending -> inProgress -> Success [Pending]
// settleOrder settles a pending order using a third-party API.j
// Implement this method using fakes defined above.
func (s *Service) settleOrder(order Order) error {
	if err := s.repo.Transact(func(tx *sql.Tx) error{
		if s.repo.FindByID(order.ID).Status == StatusPending {
			return s.repo.MarkInProgress()	
		}
		return errors.New("The order have been taken by another node")
	}); err != nil {
		return err
	}

	if _, err := s.api.Settle(order); err != nil {
		return s.repo.MarkBroken(err.Error())} ); err != nil {
			return err
		}
		return err
	}
	if err := s.api.MarkSuccess(/**/, order); err != nil {
		return err
	}
	return nil
}

func main() {
	repo := &fakeOrderRepository{}
	api := &fakeThirdParty{}

	service := Service{
		repo: repo,
		api:  api,
	}

	err := service.processOrders()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("processing finished")
	os.Exit(0)
}
