# drive-back

Para etiquetar la imagen que se construye con un Dockerfile, no se especifica el nombre de la imagen en el Dockerfile mismo. En su lugar, se proporciona como un argumento al comando `docker build` que se utiliza para construir la imagen. Aquí tienes un ejemplo de cómo podrías hacerlo:

```bash
docker build -t my-image-name:tag .
```

En este comando, `my-image-name` es el nombre que quieres darle a tu imagen y `tag` es la etiqueta que quieres usar. El punto al final del comando indica a Docker que debe buscar el Dockerfile en el directorio actual.

Por lo tanto, si quisieras construir tu imagen y llamarla `my-wompi-app` con la etiqueta `v1.0`, usarías el siguiente comando:

```bash
docker build -t my-wompi-app:v1.0 .
```

Este comando construirá la imagen utilizando el Dockerfile en el directorio actual, y luego etiquetará la imagen resultante con el nombre `my-wompi-app` y la etiqueta `v1.0`.