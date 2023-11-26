package stringList

import (
	"errors"
	"myModules/str_list_hw/node"
	"strings"
)

// StringList представляет односвязный список, представляющий строку
type StringList struct {
	Head *node.Node
}

// New создает новый StringList из обычной строки Go
func New(s string) *StringList {
	var list StringList
	for _, char := range s {
		list.Append(char)
	}
	return &list
}

// Len возвращает длину строки
func (s *StringList) Len() int {
	count := 0
	current := s.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

// Append добавляет символ в конец строки
func (s *StringList) Append(c rune) {
	newNode := &node.Node{Data: c, Next: nil}
	if s.Head == nil {
		s.Head = newNode
		return
	}

	current := s.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Prepend добавляет символ в начало строки
func (s *StringList) Prepend(c rune) {
	newNode := &node.Node{Data: c, Next: s.Head}
	s.Head = newNode
}

// Insert вставляет символ на определенную позицию в строке
func (s *StringList) Insert(c rune, index int) error {
	if index < 0 || index > s.Len() {
		return errors.New("Index out of range")
	}

	newNode := &node.Node{Data: c, Next: nil}

	if index == 0 {
		newNode.Next = s.Head
		s.Head = newNode
		return nil
	}

	current := s.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	return nil
}

// Remove удаляет символ из строки по индексу
func (s *StringList) Remove(index int) error {
	if index < 0 || index >= s.Len() {
		return errors.New("Index out of range")
	}

	if index == 0 {
		s.Head = s.Head.Next
		return nil
	}

	current := s.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	return nil
}

// String возвращает строковое представление списка
func (s *StringList) String() string {
	var builder strings.Builder
	current := s.Head
	for current != nil {
		builder.WriteRune(current.Data)
		current = current.Next
	}
	return builder.String()
}

// At получает символ по индексу
func (s *StringList) At(index int) (rune, error) {
	if index < 0 || index >= s.Len() {
		return 0, errors.New("Index out of range")
	}

	current := s.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Data, nil
}

// Concat конкатенирует две строки
func (s *StringList) Concat(other *StringList) *StringList {
	result := New("")
	current := s.Head
	for current != nil {
		result.Append(current.Data)
		current = current.Next
	}

	current = other.Head
	for current != nil {
		result.Append(current.Data)
		current = current.Next
	}

	return result
}

// Equals проверяет, равны ли две строки
func (s *StringList) Equals(other *StringList) bool {
	if s.Len() != other.Len() {
		return false
	}

	currentS := s.Head
	currentOther := other.Head

	for currentS != nil && currentOther != nil {
		if currentS.Data != currentOther.Data {
			return false
		}
		currentS = currentS.Next
		currentOther = currentOther.Next
	}

	return true
}

// IndexOf возвращает первый индекс указанного символа в строке
func (s *StringList) IndexOf(c rune) int {
	index := 0
	current := s.Head

	for current != nil {
		if current.Data == c {
			return index
		}
		index++
		current = current.Next
	}

	return -1
}

// Substring возвращает подстроку из строки
func (s *StringList) Substring(start, end int) (*StringList, error) {
	if start < 0 || start >= s.Len() || end < 0 || end > s.Len() || start > end {
		return nil, errors.New("Invalid start or end index")
	}

	substring := New("")
	current := s.Head

	for i := 0; i < start; i++ {
		current = current.Next
	}

	for i := start; i <= end; i++ {
		substring.Append(current.Data)
		current = current.Next
	}

	return substring, nil
}

// Replace заменяет символы в строке
func (s *StringList) Replace(old, new rune, n int) *StringList {
	count := 0
	current := s.Head

	for current != nil && count < n {
		if current.Data == old {
			current.Data = new
			count++
		}
		current = current.Next
	}

	return s
}
