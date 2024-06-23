package main

import (
	"errors"
	"fmt"
	"strings"
)

//4. Diccionario de contactos

type Contact struct {
	Name		string
	Phone		string
	Email 		string
	Address		string
}

type ContactList struct {
	contacts []Contact
}

func (cl *ContactList) AddContact(name, phone, email, address string){
	contact := Contact{
		Name: name,
		Phone: phone,
		Email: email,
		Address: address,
	}
	cl.contacts = append(cl.contacts, contact)
}

func (cl *ContactList) ShowContactList(){
	fmt.Println("Contact List:")
	fmt.Println("==========================")
	for i, contact := range cl.contacts{
		fmt.Println(i)
		fmt.Printf(" Name: %s\n Phone: %s\n Email: %s\n Address: %s\n", contact.Name, contact.Phone, contact.Email, contact.Address)
		fmt.Println("-------------------------")

	}
	
}

func (cl *ContactList) FindContact(query string) []Contact {
	var foundContacts []Contact
	for _, contact := range cl.contacts {
		if contact.Name == query || contact.Phone == query {
			foundContacts = append(foundContacts, contact)
			
		}
	}
	return foundContacts
}


func (cl *ContactList) UpdateContact(name string, phone, email, address *string) error {

	for i := range cl.contacts {
		if strings.EqualFold(cl.contacts[i].Name, name){

			if phone != nil {
				cl.contacts[i].Phone = *phone
			}
			if email != nil {
				cl.contacts[i].Email = *email
			}
			if address != nil {
				cl.contacts[i].Address = *address
			}
			return nil

		}

		
	}
	return errors.New("contacto no encontrado")

}
	
func (cl *ContactList) RemoveContact(name string) error {
	for i, contact := range cl.contacts {
		if contact.Name == name  {
			cl.contacts = append(cl.contacts[:i], cl.contacts[i+1:]...)
			return nil
		}
	}
	return errors.New("contacto no encontrado")	
}

func main(){
	contactList := &ContactList{}

	contactList.AddContact("Raul Gonzalez", "999222333", "raul@correo.com", "Av. Felicidad 123")
	contactList.AddContact("Carla Acuña", "999222334", "carla@correo.com", "Av. San Borja 123")
	contactList.AddContact("Otro contacto", "999222555", "otro@correo.com", "Av. San Borja 123")

	contactList.ShowContactList()

	fmt.Println("Búsqueda")

	//busqueda por nombre
	foundContacts := contactList.FindContact("Raul Gonzalez")
	for _, contact := range foundContacts{
		fmt.Printf(" Name: %s\n Phone: %s\n Email: %s\n Address: %s\n", contact.Name, contact.Phone, contact.Email, contact.Address)
	}
	fmt.Println("*************************")
	
	//busqueda por teléfono
	foundContacts = contactList.FindContact("999222334")
	for _, contact := range foundContacts{
		fmt.Printf(" Name: %s\n Phone: %s\n Email: %s\n Address: %s\n", contact.Name, contact.Phone, contact.Email, contact.Address)
	}

	//actualiza contacto
	
	newEmail := "rg@mail.com"
	err := contactList.UpdateContact("Raul Gonzalez", nil, &newEmail, nil)
	if err != nil{
		fmt.Println("Error:", err)
	}
	contactList.ShowContactList()

	//elimina contacto
	err = contactList.RemoveContact("Otro contacto")
	if err != nil{
		fmt.Println("Error:", err)
	}
	contactList.ShowContactList()
	
}




/*
//3. Software bibliotecario

type BookStatus string

const (
	Available BookStatus = "disponible"
	Borrowed  BookStatus = "prestado"
)

type Book struct {
	Title  string
	Author string
	Genre  string
	Status BookStatus
}

type Library struct {
	books []Book
}

func (l *Library) FindBooks(query string) []Book {
	var foundBooks []Book
	for _, book := range l.books {
		if book.Title == query || book.Author == query {
			foundBooks = append(foundBooks, book)
		}
	}
	return foundBooks
}

func (l *Library) AddBook(title, author, genre string) {
	book := Book{
		Title:  title,
		Author: author,
		Genre:  genre,
		Status: Available,
	}
	l.books = append(l.books, book)
}

func (l *Library) UpdateBookStatus(title string, status BookStatus) error {
	for i, book := range l.books {
		if book.Title == title {
			l.books[i].Status = status
			return nil
		}
	}
	return errors.New("Libro no encontrado")
}

func (l *Library) RemoveBook(title string) error {
	for i, book := range l.books {
		if book.Title == title {
			l.books = append(l.books[:i], l.books[i+1:]...)
			return nil
		}
	}
	return errors.New("libro no encontrado")
}

func main() {
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")

	err := library.UpdateBookStatus("El Quijote", Borrowed)
	if err != nil {
		fmt.Println("Error actualizando el estado del libro: ", err)
	}

	foundBooks := library.FindBooks("Gabriel García Márquez")
	fmt.Println("Libros encontrados: ")
	for _, book := range foundBooks {
		fmt.Printf("Título: %s, Autor: %s, Género: %s, Estado: %s\n", book.Title, book.Author, book.Genre, book.Status)
	}

	err = library.RemoveBook("Cien años de soledad")
	if err != nil {
		fmt.Println("Error eliminando el libro: ", err)
	}

	fmt.Println("Inventario completo: ")
	for _, book := range library.books {
		fmt.Printf("Título: %s, Autor: %s, Género: %s, Estado: %s\n", book.Title, book.Author, book.Genre, book.Status)
	}

}
*/
/*
//2.Sistema de gestión de tareas

import (
	"errors"
	"fmt"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pendiente"
	InProgress TaskStatus = "en progreso"
	Completed  TaskStatus = "completada"
)

type Task struct {
	Name        string
	Description string
	Assignee    string
	Status      TaskStatus
}

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(name, description string) {
	task := Task{

		Name:        name,
		Description: description,
		Status:      Pending,
	}
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) AssignTask(name, assignee string) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Assignee = assignee
			return nil
		}
	}
	return errors.New("tarea no encontrada")
}

func (tm *TaskManager) UpdateTaskStatus(name string, status TaskStatus) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Status = status
			return nil
		}
	}
	return errors.New("Tarea no encontrada")
}

func (tm *TaskManager) PendingTasks() []Task {
	var pendingTasks []Task
	for _, task := range tm.tasks {
		if task.Status == Pending {
			pendingTasks = append(pendingTasks, task)
		}
	}
	return pendingTasks
}

func main() {
	tm := &TaskManager{}

	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripcion de la tarea 2")

	err := tm.AssignTask("Tarea 1", "Juan")
	if err != nil {
		fmt.Println("Error actualizando estado de tarea: ", err)
	}

	pendingTasks := tm.PendingTasks()
	fmt.Printf("Tareas pendientes: ")
	for _, task := range pendingTasks {
		fmt.Printf("Nombre: %s, Descripcion: %s, Responsable: %s, Estado: %s\n", task.Name, task.Description, task.Assignee, task.Status)
	}
}
*/
/*
//1.Gestion de Inventario

type Producto struct {
	Nombre		string
	Precio  	float32
	Cantidad	int
}

type Inventario struct {
	productos []Producto
}



func (inv *Inventario) CrearProducto(nombre string, precio float32, cantidad int) {
	nuevoProducto:= Producto{
		Nombre:		nombre,
		Precio:		precio,
		Cantidad:	cantidad,
	}
	inv.productos = append(inv.productos, nuevoProducto)

}

func (inv *Inventario) ActualizarProducto(nombre string, cantidad int) error {
	for i, producto:= range inv.productos{
		if producto.Nombre == nombre {
			inv.productos[i].Cantidad = cantidad
			return nil
		}
	}
	return errors.New("Producto no encontrado")
}

func (inv *Inventario) EliminarProducto(nombre string) error {
	for i, producto:= range inv.productos{
		if producto.Nombre == nombre {
			inv.productos = append(inv.productos[:i], inv.productos[i+1:]... )
			return nil
		}
	}
	return errors.New("Producto no encontrado")
}

func (inv *Inventario) MostrarInventario() {
	for _, producto :=range inv.productos {
		fmt.Printf(" Nombre: %s\n Precio: %.2f\n Cantidad: %d\n", producto.Nombre, producto.Precio, producto.Cantidad)
		fmt.Println("--------------------")
	}
}

func main(){
	inventario := &Inventario{}

	//agregando nuevos productos al inventario
	inventario.CrearProducto("TV Samsung 32 pulgadas", 1049.99, 10)
    inventario.CrearProducto("Monitor LG 48 pulgadas",3500, 5)
	inventario.CrearProducto("Teclado LG",450.50, 100)

    //mostrar inventario
	inventario.MostrarInventario()

	fmt.Println("=====================================")

	//actualizar Producto
	err := inventario.ActualizarProducto("Teclado LG",50)
	if err != nil{
		fmt.Println("Error al actualizar la cantidad del producto:",err)
	}
	inventario.MostrarInventario()


	fmt.Println("=====================================")

	//eliminar un producto
	err = inventario.EliminarProducto("Teclado LG")
	if err != nil{
		fmt.Println("Error al eliminar el producto:",err)
	}

	inventario.MostrarInventario()

}
*/