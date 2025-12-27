## 1. Actualización de repositorios e instalación de paquetes necesarios
```bash
sudo apt update
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common
```

## 2. Añadir la clave gpg oficial de Docker

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

## 3. Añadir el repositorio Docker a las fuentes de APT

```bash
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
```

## 4. Actualizar la base de datos de paquetes e instalar DockerCE y Docker-Compose

```bash
sudo apt update

sudo apt install -y docker-ce docker-ce-cli containerd.io

sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```


## 5. Añadir tu usuario al grupo docker para evitar usar sudo

```bash
sudo usermod -aG docker $USER
```

## 6. Verificar instalaciones
```bash
docker --version
docker-compose --version
```

&nbsp;

> Después de ejecutar estos comandos, es recomendable cerrar la terminal y volver a iniciarla de modo que se apliquen todos los cambios.