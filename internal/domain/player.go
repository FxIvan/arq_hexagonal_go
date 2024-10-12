/**********************
*El paquete domain contiene las entidades principales y la lógica del dominio de la aplicación.
*Aquí es donde se definen los objetos centrales del negocio, como Player, y su comportamiento.
*Este paquete es completamente independiente de cualquier infraestructura, como bases de datos
*o frameworks externos, lo que permite que el núcleo del dominio sea reutilizable y testeable de manera aislada.
*
*Player representa una entidad del dominio que encapsula los atributos clave de un jugador,
*como su nombre, edad y el momento en que fue creado dentro del sistema. Los campos están etiquetados
*para su uso con JSON y reglas de validación.
*package domain
*********************/
package domain

import "time"

type Player struct {
	Name         string    `json:"name"	binding:"required"`
	Age          int       `json:"age"		binding:"required"`
	CreationTime time.Time `json:"creation_time"`
}
