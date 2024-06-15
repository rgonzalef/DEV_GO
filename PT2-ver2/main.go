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
	for _, contact := range cl.contacts{
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
	return errors.New("Contacto no encontrado")

}
	
func (cl *ContactList) RemoveContact(name string) error {
	for i, contact := range cl.contacts {
		if contact.Name == name  {
			cl.contacts = append(cl.contacts[:i], cl.contacts[i+1:]...)
			return nil
		}
	}
	return errors.New("Contacto no encontrado")	
}

func main(){
	contactList := ContactList{}

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
//3.Software bibliotecario

type BookStatus string

const (
	Available BookStatus = "disponible"
	Borrowed  BookStatus = "prestado"
)

type Book struct {
	Title string
	Author  string
	Genre string
	Status BookStatus
}

type Library struct {
	books []Book
}

func (l *Library) AddBook(title, author, genre string) {
	book := Book{
		Title: title,
		Author: author,
		Genre: genre,
		Status: Available,
	}
	l.books = append(l.books, book)
}

func (l *Library) showLibrary(){
	fmt.Println("Colección de libros de la Libreria:")
	fmt.Println("==========================")
	for _, book := range l.books {
		fmt.Printf(" Titulo: %s\n Autor: %s\n Género: %s\n Status: %s\n",book.Title, book.Author, book.Genre, book.Status)
		fmt.Println("---------------------------")
	}
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

func (l *Library) FindBooks(query string) []Book {
	var foundBooks []Book
	for _, book := range l.books {
		if book.Title == query || book.Author == query {
			foundBooks = append(foundBooks, book)
		}
	}
	return foundBooks
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

func main(){
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")
	library.AddBook("Inferno", "Dan Brown", "Suspenso")
	library.showLibrary()

	library.UpdateBookStatus("Inferno", Borrowed)
	library.UpdateBookStatus("El Quijote", Borrowed)
	library.showLibrary()

	foundBooks := library.FindBooks("Dan Brown")
	fmt.Println("Libros encontrados: ")
	for _, book := range foundBooks {
		fmt.Printf("Título: %s, Autor: %s, Género: %s, Estado: %s\n", book.Title, book.Author, book.Genre, book.Status)
	}
	foundBooks= library.FindBooks("El Quijote")
	fmt.Println("Libros encontrados: ")
	for _, book := range foundBooks {
		fmt.Printf("Título: %s, Autor: %s, Género: %s, Estado: %s\n", book.Title, book.Author, book.Genre, book.Status)
	}

	err := library.RemoveBook("Cien años de soledad")
	if err != nil {
		fmt.Println("Error eliminando el libro: ", err)
	}
	library.showLibrary()
}
*/

/*
//2.Sistema de gestión de tareas

type TaskStatus string

const (
	Pending    TaskStatus = "pendiente"
	InProgress TaskStatus = "en progreso"
	Completed  TaskStatus = "completada"
)

type Task struct {
	Name        string
	Description string
	Responsible    string
	Status      TaskStatus
}

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(name, description string) {
	task := Task{
		Name:			name,
		Description: 	description,
		Status:			Pending,
	}
	tm.tasks = append(tm.tasks, task)

}

func (tm *TaskManager) printTaskManager(){
	for _, task :=range tm.tasks {
		fmt.Printf(" Name: %s\n Description: %s\n Responsible: %s\n Status: %s\n", task.Name, task.Description, task.Responsible, task.Status)
		fmt.Println("--------------------")
	}
}

func (tm *TaskManager) AssignTask(name, responsible string) error {
	for i, task := range tm.tasks {
		if task.Name == name {
			tm.tasks[i].Responsible = responsible
			return nil
		}
	}
	return errors.New("Tarea no encontrada")
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

	//creando nuevas tareas
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripcion de la tarea 2")
	tm.AddTask("Tarea 3", "Descripcion de la tarea 3")

	tm.printTaskManager()
	fmt.Println("=================================")

	err := tm.AssignTask("Tarea 1", "Raul Gonzalez")
	if err != nil{
		fmt.Println("Error actualizando estado de tarea: ", err)
	}
	tm.AssignTask("Tarea 2", "Sami Gonzalez")
	if err != nil{
		fmt.Println("Error actualizando estado de tarea: ", err)
	}

	tm.printTaskManager()
	fmt.Println("=================================")

	tm.UpdateTaskStatus("Tarea 1", InProgress)
	tm.UpdateTaskStatus("Tarea 2", Completed)
	tm.printTaskManager()
	fmt.Println("=================================")

	pendingTasks := tm.PendingTasks()
	fmt.Println("Tareas pendientes: ")
	for _, task := range pendingTasks {
		fmt.Printf("Nombre: %s\n, Descripcion: %s\n, Responsable: %s\n, Estado: %s\n", task.Name, task.Description, task.Responsible, task.Status)
	}
}
*/

///////////////////////////////////////////////////////////////////////////////////////
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



func (inv *Inventario) crearProducto(nombre string, precio float32, cantidad int) {
	nuevoProducto:= Producto{
		Nombre:		nombre,
		Precio:		precio,
		Cantidad:	cantidad,
	}
	inv.productos = append(inv.productos, nuevoProducto)

}

func (inv *Inventario) actualizarProducto(nombre string, cantidad int) error {
	for i, producto:= range inv.productos{
		if producto.Nombre == nombre {
			inv.productos[i].Cantidad = cantidad
			return nil
		}
	}
	return errors.New("Producto no encontrado")
}

func (inv *Inventario) eliminarProducto(nombre string) error {
	for i, producto:= range inv.productos{
		if producto.Nombre == nombre {
			inv.productos = append(inv.productos[:i], inv.productos[i+1:]... )
			return nil
		}
	}
	return errors.New("Producto no encontrado")
}

func (inv *Inventario) mostrarInventario() {
	for _, producto :=range inv.productos {
		fmt.Printf(" Nombre: %s\n Precio: %.2f\n Cantidad: %d\n", producto.Nombre, producto.Precio, producto.Cantidad)
		fmt.Println("--------------------")
	}
}

func main(){
	inventario := &Inventario{}

	//agregando nuevos productos al inventario
	inventario.crearProducto("TV Samsung 32 pulgadas", 1049.99, 10)
	inventario.crearProducto("Monitor LG 48 pulgadas",3500, 5)
	inventario.crearProducto("Teclado LG",450.50, 100)

	//mostrar inventario
	inventario.mostrarInventario()

	fmt.Println("=====================================")

	//actualizar Producto
	err := inventario.actualizarProducto("Teclado LG",50)
	if err != nil{
		fmt.Println("Error al actualizar la cantidad del producto:",err)
	}
	inventario.mostrarInventario()


	fmt.Println("=====================================")

	//eliminar un producto
	err = inventario.eliminarProducto("Teclado LG")
	if err != nil{
		fmt.Println("Error al eliminar el producto:",err)
	}

	inventario.mostrarInventario()

}
*/