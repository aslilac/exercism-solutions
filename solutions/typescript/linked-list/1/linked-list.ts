export class LinkedList<T> {
  head: Node<T> | null = null;
  tail: Node<T> | null = null;
  
  public push(value: T) {
    const node = new Node(value);
    if (!this.head) {
      this.head = node;
    }
    if (this.tail) {
      node.prev = this.tail;
      this.tail.next = node;
    }
    this.tail = node;
  }

  public pop(): T | undefined {
    const node = this.tail;
    if (!node) {
      return undefined;
    }
    
    if (node.prev) {
      this.tail = node.prev;
      this.tail.next = null;
    } else {
      this.head = null;
      this.tail = null;
    }

    return node.value;
  }

  public shift(): T | undefined {
    const node = this.head;
    if (!node) {
      return undefined;
    }

    if (node.next) {
      this.head = node.next;
      this.head.prev = null;
    } else {
      this.head = null;
      this.tail = null;
    }

    return node.value;
  }

  public unshift(value: T) {
    const node = new Node(value);
    if (!this.tail) {
      this.tail = node;
    }
    if (this.head) {
      node.next = this.head;
      this.head.prev = node;
    }
    this.head = node;
  }

  public delete(value: T) {
    let cursor = this.head;

    while (cursor) {
      if (cursor.value !== value) {
        cursor = cursor.next;
        continue;
      }

      if (cursor.next && cursor.prev) {
        cursor.prev.next = cursor.next;
        cursor.next.prev = cursor.prev;
        return;
      }

      if (!cursor.next && !cursor.prev) {
        this.head = null;
        this.tail = null;
        return;
      }

      if (!cursor.prev) {
        this.head = cursor.next!;
        this.head.prev = null;
      }

      if (!cursor.next) {
        this.tail = cursor.prev!;
        this.tail.next = null
      }
      
      return;
    }
  }

  public count(): number {
    let count = 0;
    let cursor = this.head;

    while (cursor) {
      count++;
      cursor = cursor.next;
    }

    return count;
  }
}

class Node<T> {
  prev: Node<T> | null = null;
  next: Node<T> | null = null;

  constructor(readonly value: T) {}
}
