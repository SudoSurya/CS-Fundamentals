class Node {
    int data;
    Node next;
    Node prev;

    Node(int data) {
        this.data = data;
        this.next = null;
        this.prev = null;
    }
}

public class DoubleLinkedList {
    public Node prev;
    public Node next;

    public DoubleLinkedList(Node prev, Node next) {
        this.prev = prev;
        this.next = next;
    }

    public void insert_at_end(int value) {
        Node new_node = new Node(value);

        if (prev == null) {
            prev = new_node;
            next = new_node;
            return;
        } else {
            new_node.prev = next;
            next.next = new_node;
            next = new_node;
        }
    }

    public void display() {
        Node current = prev;
        while (current != null) {
            System.out.println(current.data);
            current = current.next;
        }
    }

    public static void main(String[] args) {
        DoubleLinkedList list = new DoubleLinkedList(null, null);
        list.insert_at_end(1);
        list.insert_at_end(2);
        list.insert_at_end(3);
        // rrint list

    }

}
