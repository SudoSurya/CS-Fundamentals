class InnerQueueImpl {
    int[] queue;

    public InnerQueueImpl(int size) {
        queue = new int[size];
    }

    public void enqueue(int value) {
        for (int i = 0; i < queue.length; i++) {
            if (queue[i] == 0) {
                queue[i] = value;
                break;
            }
        }
    }

    public int dequeue() {
        int value = queue[0];
        for (int i = 0; i < queue.length - 1; i++) {
            queue[i] = queue[i + 1];
        }
        queue[queue.length - 1] = 0;
        return value;
    }

    public int peek() {
        return queue[0];
    }

}

public class QueueImpl {

    public static void main(String[] args) {

        InnerQueueImpl queue = new InnerQueueImpl(8);
        queue.enqueue(1);
        queue.enqueue(2);
        System.out.println(queue.peek());
    }

}
