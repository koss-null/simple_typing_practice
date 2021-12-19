# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        cur = head
        length = 0
        while cur != None:
            length += 1
            cur = cur.next

        cur = head
        prev = None
        place = 0
        while cur != None:
            if place == length - n:
                prev.next = cur.next
                cur = prev
            prev = cur
            cur = cur.next
            place += 1
        return head
