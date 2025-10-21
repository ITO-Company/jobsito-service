# Jobsito

## Profile module

### Job seeker profiles (postulantes)

#### Feature(high): Config

Se necesita configurar la conexión con la db, crear todos los modelos y usar semillas para el gobal tag.

#### Feature(high): Auth

Se requiere integrar las autenticaciones al jobseeker o postulante y a la company o empresa, es decir, que se pueda hacer un signup y signin para el postulante y la empresa, se debe configurar a la par el jwt, junto al middleware de autenticación, hacer lo mismo con el middleware de permisos, teniendo en cuenta que tendremos permisos y roles estáticos, como los son los postulantes y las compañias.

#### Feature(mid): Jobseeker profile (jobseeker)

Hacer endpoints para que el postulante, pueda ver su perfil (cuando ve su perfil, se necesita ver todas las postulaciones que realizó), editar su perfil y borrar su perfil, si es que ya no quiere estar en la aplicación, se debe serciorar que sea el postulante único que ocupe estos endpoints usando el jwt para esto.

#### Feature(mid): Search jobseekers (company)

Se requiere un endpoint donde se enliste a los portulantes, con ciertos criterios de búsqueda (se pueden usar las jobseeker tags) y usando paginación, este endpoint es exclusivo de las empresas

#### Feature(low): View profile (company)

Se requiere que un endpoint para la compañia, donde esta pueda ver el perfil del postulante (sin las postulaciones de este último).

### Company profiles (empresa)

#### Feature(high): Profile (company)

Se requieren endpoints para editar, borrar y visualizar (se tiene que obtener también todas las ofertas laborales ofertadas por la empresa) el perfil de una empresa, estos endpoints son solo para una empresa, se debe evaluar que sea la empresa única usando el jwt.

#### Feature(mid): Search Companies (jobseeker)

Se necesita que el postulante pueda buscar compañías usando criterios de búsqueda (como el company tag) y paginación.

#### Feature(low): Company profile (jobseeker)

El postulante puede acceder a la vista del perfilde la empresa, más solo ver su perfil, no puede ver otra cosa.

### Global tags

#### Feature(mid): Search global tags (both)

Hacer un endpoint con paginación y criterios de búsqueda para las etiquetas globales.

#### Feature(low): Assign jobseeker tag (jobseeker)

El postulante puede asignarse tags globales y colocarlas en su perfil por medio de la tabla jobseeker tag, también puede quitarse las tags asignadas.a

## Offers module

### JOB_POSTINGS

#### Feature(high): Post a job

Se necesita que la empresa pueda postear un empleo y editarlo, a su vez a la hora de crear o editar dicho empleo, se pueden colocar tags al empleo, usando job posting tags, además la empresa puede borrar las ofertas (eliminado lógico).

#### Feature(high): Search a job

Un postulante puede buscar una oferta laboral por medio de criterios de búsqueda y paginación.

#### Feature(mid): View a job

Un postulante necesita ver una oferta laboral, junto a las tags asociadas a dicha oferta.

## Application module

### APPLICATIONS y APPLICATION_STATUS_HISTORY

#### Feature(high): Apply (jobseeker)

Se requiere un endpoint para postularse a una oferta laboral, rellenando los datos correspondientes.

#### Feature(high): View offers apply (company)

Se requiere un endpoint que traiga todas las postulaciones a una oferta laboral en concreto

#### Feature(mid): Apply status (jobseeker)

Se requiere un endpoint para ver el estado de la aplicación, se debe usar application status history

#### Feature(mid): Manage applications (company)

Se requiere un endpoint para aceptar o rechazar una postulación, si se acepta una postulación se debe cambiar el estado de la oferta a cerrada.

### Saved jobs

#### Feature(low): Manage saved jobs (jobseeker)

Se requiere un endpoint para guardar una oferta laboral, y para quitarla de las ofertas guardadas, además de poder visualizar todas las ofertas.
