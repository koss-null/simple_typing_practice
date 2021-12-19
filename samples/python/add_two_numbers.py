# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        mem = 0
        result_head = ListNode()
        result_cur = result_head
        result_prev = None
        num1, num2 = l1, l2

        while num1 or num2:
            n1val = num1.val if num1 else 0
            n2val = num2.val if num2 else 0
            digit_sum = n1val + n2val + mem
            first_digit = digit_sum % 10
            mem = int(digit_sum / 10)
            result_cur.val = first_digit
            result_cur.next = ListNode()
            num1, num2 = num1.next, num2.next
            result_prev = result_cur
            result_cur = result_cur.next
        if mem:
            result_cur.val = mem
            result_prev = result_cur.next
        result_prev.next = None

        return result_head


