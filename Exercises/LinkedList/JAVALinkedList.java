class Node {
    public int data;
    public Node next;

    public Node(int data) {
        this.data = data;
        this.next = null;
    }
}

public class JAVALinkedList {
    public Node head;

    public JAVALinkedList(Node head) {
        this.head = null;
    }

    public int Read(int index) {
        Node current = head;
        int currentIndex = 0;

        while (currentIndex < index) {
            current = current.next;
            currentIndex++;
            if (current == null) {
                return -1;
            }
        }
        return current.data;
    }

    public void deleteNode(int index) {
        if (index == 0) {
            this.head = this.head.next;
            return;
        }
        Node current = this.head;
        int currentIndex = 0;
        while (currentIndex < index - 1) {
            current = current.next;
            currentIndex++;
        }
        Node nextNode = current.next.next;
        current.next = nextNode;
    }

    public static void main(String[] args) {
        JAVALinkedList list = new JAVALinkedList(null);
        list.head = new Node(1);
        Node second = new Node(2);
        Node third = new Node(3);

        list.head.next = second;
        second.next = third;

        System.out.println(list.Read(1));

    }
}
