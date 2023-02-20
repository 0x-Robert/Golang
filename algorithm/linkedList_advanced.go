package main

import "fmt"

//노드 구조체, 데이터와 노드로 이루어짐
type Node struct {
    data int
    next *Node
}

//링크드 리스트 구조체 
//노드의 헤드가 있음 
type LinkedList struct {
    head *Node
}

//노드를 추가하는 함수 
//AddNode 메서드는 지정된 데이터가 있는 새 노드를 목록 끝에 추가합니다. 
//먼저 주어진 데이터로 새 노드를 만든 다음 목록이 비어 있는지 확인합니다(즉, 헤드가 0인지 여부)
//이 경우 새 노드가 목록의 머리글로 설정됩니다.
// 그렇지 않으면 마지막 노드에 도달할 때까지 목록을 반복한 다음 마지막 노드의 다음 필드가 새 노드를 가리키도록 설정합니다.
func (list *LinkedList) AddNode(data int) {
    newNode := &Node{data, nil}

    if list.head == nil {
        list.head = newNode
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNode
    }
}

//RemoveNode 메서드는 지정된 데이터가 있는 목록의 첫 번째 노드를 제거합니다. 
//목록이 비어 있으면 단순히 반환됩니다. 
//목록의 첫 번째 노드에 주어진 데이터가 있으면 목록의 머리가 두 번째 노드를 가리키도록 설정합니다(두 번째 노드가 없으면 0으로 설정됩니다). 
//그렇지 않으면 지정된 데이터가 있는 노드를 찾을 때까지 목록을 반복한 다음 제거된 노드를 가리키도록 이전 노드의 다음 필드를 설정합니다.

func (list *LinkedList) RemoveNode(data int) {
    if list.head == nil {
        return
    }

    if list.head.data == data {
        list.head = list.head.next
        return
    }

    previous := list.head
    for previous.next != nil {
        if previous.next.data == data {
            previous.next = previous.next.next
            return
        }
        previous = previous.next
    }
}

// 단순히 목록을 반복하고 각 노드의 데이터를 인쇄합니다.
func (list *LinkedList) PrintList() {
    current := list.head
    for current != nil {
        fmt.Print(current.data, " ")
        current = current.next
    }
    fmt.Println()
}

func main() {
 
	// 이 예에서는 새 LinkedList 인스턴스를 생성하고 여기에 세 개의 노드를 추가한 다음 목록을 인쇄하고 
	//목록에서 두 번째 노드를 제거한 다음 목록을 다시 인쇄합니다. 입력 및 출력은 다음과 같습니다:
	myList := &LinkedList{}
    myList.AddNode(1)
    myList.AddNode(2)
    myList.AddNode(3)
    myList.PrintList()
    myList.RemoveNode(2)
    myList.PrintList()
}