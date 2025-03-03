# Actions

Las acciones tienen como objetivo manejar las peticiones
segun el tipo de caso que requieran, ejemplo:
    
> si la peticion esta hecha para que borre un registro
de los animes, se le concedera la peticion de borrar los
registros a las acciones de anime, donde mandara realizar
el cambio a la DB.

Por este motivo es que los verbos estan separados de las acciones, de este modo se pueden realizar cualquier tipo de peticion al servidor y en base a la ruta y el cuerpo de la peticion es como actuara el sistema de acciones.