package main

import (
	"testing"
)

//4. Diccionario de contactos -TESTING

func TestAddContact(t *testing.T){
	cl := &ContactList{}
	cl.AddContact("Raul Gonzalez", "999222333", "raul@correo.com", "Av. Felicidad 123")
	

	if len(cl.contacts) != 1 {
		t.Errorf("Se esperaba un contacto, se obtuvo %d", len(cl.contacts))
	}

	if cl.contacts[0].Name != "Raul Gonzalez" {
		t.Errorf("Contacto incorrecto se esperaba 'Raul Gonzalez', se obtuvo %s", cl.contacts[0].Name )
	}
}

func TestFindContact(t *testing.T){
	cl := &ContactList{}
	cl.AddContact("Raul Gonzalez", "999222333", "raul@correo.com", "Av. Felicidad 123")

	fc := cl.FindContact("Raul Gonzalez")

	if len(fc) != 1 {
		t.Errorf("Se esperaba un contacto, se obtuvo %d", len(fc))
	}

	if fc[0].Name != "Raul Gonzalez"{
		t.Errorf("Se esperaba el contacto 'Raul Gonzalez', se obtuvo %s", fc[0].Name )
	}
}

func TestUpdateContact(t *testing.T){
	cl := &ContactList{}
	cl.AddContact("Raul Gonzalez", "999222333", "raul@correo.com", "Av. Felicidad 123")

	newEmail := "rg@mail.com"
	err := cl.UpdateContact("Raul Gonzalez", nil, &newEmail, nil)

	if err != nil {
		t.Errorf("Error actualizando el contacto: %v", err)
	}

	if cl.contacts[0].Email != newEmail{
		t.Errorf("Campo email no actualizado correctamente, se esperaba %s y se obtuvo %s", newEmail, cl.contacts[0].Email )
	}
}

func TestRemoveContact(t *testing.T){
	cl := &ContactList{}

	cl.AddContact("Raul Gonzalez", "999222333", "raul@correo.com", "Av. Felicidad 123")
	cl.AddContact("Carla Acuña", "999222334", "carla@correo.com", "Av. San Borja 123")

	err := cl.RemoveContact("Carla Acuña")

	if err != nil {
		t.Errorf("Error eliminando contacto %v", err)
	}

	if len(cl.contacts) != 1 {
		t.Errorf("Se esperaba 1 contacto, se obtuvo %d", len(cl.contacts) )
	}
}

func TestContactIntegration(t *testing.T){

	cl := &ContactList{}

	cl.AddContact("Raul Gonzalez", "999222333", "raul@correo.com", "Av. Felicidad 123")
	cl.AddContact("Carla Acuña", "999222334", "carla@correo.com", "Av. San Borja 123")

	fc := cl.FindContact("Raul Gonzalez")

	if fc[0].Name != "Raul Gonzalez"{
		t.Errorf("Se esperaba el contacto 'Raul Gonzalez', se obtuvo %s", fc[0].Name )
	}

	newEmail := "rg@mail.com"
	err := cl.UpdateContact("Raul Gonzalez", nil, &newEmail, nil)

	if err != nil {
		t.Errorf("Error actualizando el contacto: %v", err)
	}

	if cl.contacts[0].Email != newEmail{
		t.Errorf("Campo email no actualizado correctamente, se esperaba %s y se obtuvo %s", newEmail, cl.contacts[0].Email )
	}

	err = cl.RemoveContact("Carla Acuña")

	if err != nil {
		t.Errorf("Error eliminando contacto %v", err)
	}

	if len(cl.contacts) != 1 {
		t.Errorf("Se esperaba 1 contacto, se obtuvo %d", len(cl.contacts) )
	}
}

/*
//3. Software bibliotecario - TESTING
func TestAddBook(t *testing.T) {

	library := &Library{}
	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")

	foundBooks := library.FindBooks("Miguel de Cervantes")
	if len(foundBooks) != 1 {
		t.Errorf("Se esperaba 1 libro, se obtuvo %d", len(foundBooks))
	}

	if foundBooks[0].Title != "El Quijote" {
		t.Errorf("Título de libro encontrado incorrecto, se esperaba 'El Quijote', se obtuvo '%s'", foundBooks[0].Title)
	}
}

func TestUpdateBooksStatus(t *testing.T) {
	library := &Library{}
	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	err := library.UpdateBookStatus("El Quijote", Borrowed)

	if err != nil {
		t.Errorf("Error actualizando estado del libro: %v", err)
	}

	if library.books[0].Status != Borrowed {
		t.Errorf("Estado del libro incorrecto, se esperaba '%s', se obtuvo '%s'", Borrowed, library.books[0].Status)
	}
}

func TestRemoveBook(t *testing.T) {
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	err := library.RemoveBook("El Quijote")

	if err != nil {
		t.Errorf("Error eliminando libro: %v", err)
	}

	if len(library.books) != 0 {
		t.Errorf("Se esperaba 0 libros, se obtuvo %d", len(library.books))
	}

}

func TestLibraryIntegration(t *testing.T) {
	library := &Library{}

	library.AddBook("El Quijote", "Miguel de Cervantes", "Novela")
	library.AddBook("Cien años de soledad", "Gabriel García Márquez", "Novela")

	err := library.UpdateBookStatus("El Quijote", Borrowed)
	if err != nil {
		t.Errorf("Error actualizando estado del libro: %v", err)
	}

	foundBooks := library.FindBooks("Gabriel García Márquez")
	if len(foundBooks) != 1 {
		t.Errorf("Se esperaba 1 libro encontrado, se obtuvo %d", len(foundBooks))
	}

	if foundBooks[0].Title != "Cien años de soledad" {
		t.Errorf("Título del libro encontrado incorrecto, se esperaba 'Cien años de soledad', se obtuvo '%s'", foundBooks[0].Title)
	}

	err = library.RemoveBook("Cien años de soledad")
	if err != nil {
		t.Errorf("Error eliminando libro: %v", err)
	}

	if len(library.books) != 1 {
		t.Errorf("Se esperaba 1 libro en la colección, se obtuvo %d", len(library.books))
	}

	if library.books[0].Title != "El Quijote" {
		t.Errorf("Título del libro restante incorrecto, se esperaba 'El Quijote', se obtuvo '%s'", library.books[0].Title)
	}
}
*/

/*
//2.Sistema de gestión de tareas - TESTING
func TestAddTask(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripcion de la tarea 1")

	if len(tm.tasks) != 1 {
		t.Errorf("Se esperaba 1 tarea, se obtuvo %d", len(tm.tasks))
	}

	if tm.tasks[0].Name != "Tarea 1" {
		t.Errorf("Nombre de la tarea incorrecto, se esperaba 'Tarea 1', se obutuvo %s", tm.tasks[0].Name)
	}
}

func TestAssignTask(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	err := tm.AssignTask("Tarea 1", "Juan")

	if err != nil {
		t.Errorf("Error asignando tarea: %v", err)
	}

	if tm.tasks[0].Assignee != "Juan" {
		t.Errorf("Responsable incorrecto, se esperaba 'Juan', se obtuvo '%s' ", tm.tasks[0].Assignee)
	}
}

func TestUpdateTaskStatus(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	err := tm.UpdateTaskStatus("Tarea 1", InProgress)

	if err != nil {
		t.Errorf("Error actualizando estado de tarea: %v", err)
	}

	if tm.tasks[0].Status != InProgress {
		t.Errorf("Estado incorrecto, se esperaba '%s', se obtuvo '%s'", InProgress, tm.tasks[0].Status)
	}
}

func TestPendingTasks(t *testing.T) {
	tm := &TaskManager{}
	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripción de la tarea 2")
	tm.UpdateTaskStatus("Tarea 1", InProgress)

	pendingTasks := tm.PendingTasks()
	if len(pendingTasks) != 1 {
		t.Errorf("Se esperaba 1 tarea pendiente, se obtuvo %d", len(pendingTasks))
	}

	if pendingTasks[0].Name != "Tarea 2" {
		t.Errorf("Nombre de la tarea pendiente incorrecto, se esperaba 'Tarea 2', se obtuvo '%s'", pendingTasks[0].Name)
	}
}

func TestTaskManagerIntegration(t *testing.T) {
	tm := &TaskManager{}

	tm.AddTask("Tarea 1", "Descripción de la tarea 1")
	tm.AddTask("Tarea 2", "Descripción de la tarea 2")

	err := tm.AssignTask("Tarea 1", "Juan")
	if err != nil {
		t.Errorf("Error asignando tarea: %v", err)
	}

	err = tm.UpdateTaskStatus("Tarea 2", InProgress)
	if err != nil {
		t.Errorf("Error actualizando estado de tarea: %v", err)
	}

	pendingTasks := tm.PendingTasks()
	if len(pendingTasks) != 1 {
		t.Errorf("Se esperaba 1 tarea pendiente, se obtuvo %d", len(pendingTasks))
	}

	if pendingTasks[0].Name != "Tarea 1" {
		t.Errorf("Nombre de la tarea pendiente incorrecto, se esperaba 'Tarea 1', se obtuvo '%s'", pendingTasks[0].Name)
	}

	if pendingTasks[0].Assignee != "Juan" {
		t.Errorf("Responsable incorrecto, se esperaba 'Juan', se obtuvo '%s'", pendingTasks[0].Assignee)
	}
}*/

/*
//1.Gestion de Inventario-TESTING

func TestCrearProducto(t *testing.T) {
	inv := &Inventario{}
	inv.CrearProducto("TV Samsung 32 pulgadas", 1049.99, 10)

	if len(inv.productos) != 1 {
		t.Errorf("Se esperaba un registro, se obtuvo %d", len(inventario))
	}

	if inv.productos[0].Nombre != "TV Samsung 32 pulgadas" {
		t.Errorf("Producto incorrecto, se esperaba 'TV Samsung 32 pulgadas', se obtuvo '%s'", inv.productos[0].Nombre )
	}

}

func TestActualizarProducto(t *testing.T){
	inv := &Inventario{}
	inv.CrearProducto("TV Samsung 32 pulgadas", 1049.99, 10)

	err:= inv.ActualizarProducto("TV Samsung 32 pulgadas",50)

	if err != nil {
		t.Errorf("Error actualizando producto: %v", err)
	}

	if inv.productos[0].Cantidad != 50 {
		t.Errorf("Actualización incorrecta, se esperaba 50, se obutuvo %d", inv.productos[0].Cantidad)
	}
}

func TestEliminarProducto(t *testing.T){
	inv := &Inventario{}
	inv.CrearProducto("TV Samsung 32 pulgadas", 1049.99, 10)

	err := inv.EliminarProducto("TV Samsung 32 pulgadas")

	if err != nil {
		t.Errorf("Error elimininando produto: %v",err)
	}

	if len(inv.productos) != 0 {
		t.Errorf("Se esperaba 0 productos, se obtuvo %d", len(inv.productos))
	}

}

func TestProductoIntegration(t *testing.T){
	inv := &Inventario{}
	inv.CrearProducto("TV Samsung 32 pulgadas", 1049.99, 10)
	inv.CrearProducto("Monitor LG 48 pulgadas",3500, 5)
	//inv.CrearProducto("Teclado LG",300, 5)

	err:= inv.ActualizarProducto("TV Samsung 32 pulgadas",50)

	if err != nil {
		t.Errorf("Error actualizando producto: %v", err)
	}

	err = inv.EliminarProducto("TV Samsung 32 pulgadas")

	if err != nil {
		t.Errorf("Error elimininando produto: %v",err)
	}

	if len(inv.productos) != 1 {
		t.Errorf("Se esperaba 1 solo registro en el inventario, se obtuvo %d",len(inv.productos))
	}

	if inv.productos[0].Nombre != "Monitor LG 48 pulgadas" {
		t.Errorf("Nombre del producto restante incorrecto, se esperaba Monitor LG 48 pulgadas, se obtuvo %s", inv.productos[0].Nombre)
	}
}
*/