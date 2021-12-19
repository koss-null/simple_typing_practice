import math

class Shape:
    pass


class Square(Shape):
    def __init__(self, side):
        self.side = side

    def area(self):
        return self.side * self.side

    def perim(self):
        return 4 * self.side


class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return math.Pi * self.radius * self.radius

    def perim(self):
        return 2 * math.Pi * self.radius

class ShapeCalculator:
    def __init__(self):
        self.shapes = []
        self.shapes.append(Square(2))
        self.shapes.append(Circle(5))

    def output(self):
        sum_areas, sum_perims = 0, 0
        for s in self.shapes:
            sum_areas += s.area()
            sum_perims += s.perim()
        print('Full area is {}'.format(sum_areas))
        print('Full perimeter is {}'.format(sum_perims))

    def as_json(self):
        return {
            'area': sum(s.area() for s in self.shapes)
        }


ShapeCalculator().output()
