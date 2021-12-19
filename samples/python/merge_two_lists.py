# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        head = ListNode()
        cur = head
        l_head = l1
        r_head = l2
        prev = None
        while l_head or r_head:
            prev = cur
            if l_head:
                if r_head:
                    if l_head.val < r_head.val:
                        cur.val = l_head.val
                        cur.next = ListNode()
                        cur = cur.next
                        l_head = l_head.next
                    else:
                        cur.val = r_head.val
                        cur.next = ListNode()
                        cur = cur.next
                        r_head = r_head.next
                else:
                    cur.val = l_head.val
                    cur.next = ListNode()
                    cur = cur.next
                    l_head = l_head.next

            if r_head:
                if l_head:
                    if l_head.val < r_head.val:
                        cur.val = l_head.val
                        cur.next = ListNode()
                        cur = cur.next
                        l_head = l_head.next
                    else:
                        cur.val = r_head.val
                        cur.next = ListNode()
                        cur = cur.next
                        r_head = r_head.next
                else:
                    cur.val = r_head.val
                    cur.next = ListNode()
                    cur = cur.next
                    r_head = r_head.next

        if prev:
            prev.next = None
        else:
            return None

        return head



