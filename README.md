# Golang File Transfer

Ce projet est composé de deux parties principales : un serveur et un client. Ces deux parties sont contenues dans deux dossiers séparés, server et client.

## Server

Le serveur est responsable de la réception des fichiers envoyés par le client. Il écoute sur le port 8080 et accepte les connexions entrantes. Lorsqu'un fichier est reçu, il est sauvegardé dans un dossier spécifique.

Pour exécuter le serveur, naviguez vers le dossier server et exécutez le fichier principal avec la commande go run main.go.

### Configuration

Vous pouvez changer le dossier de destination des fichiers reçus en modifiant la variable `outputFolder` dans le fichier `main.go`.

```go
var outputFolder = "./reception"
```

Le server écoute par défaut sur le port 8080. Vous pouvez changer ce port à la ligne 18 du fichier `main.go`.

```go
ln, err := net.Listen("tcp", ":8080")
```

## Client

Le client est responsable de l'envoi des fichiers au serveur. Il se connecte au serveur sur localhost:8080 et envoie les fichiers contenus dans un dossier spécifique.

Pour exécuter le client, naviguez vers le dossier client et exécutez le fichier principal avec la commande go run main.go.

### Configuration

Vous pouvez changer le dossier contenant les fichiers à envoyer en modifiant la variable `payloadFolder` dans le fichier `main.go`.

```go
var payloadFolder = "./payload/"
```

Le client se connecte par défaut au serveur sur localhost:8080. Vous pouvez changer cette adresse à la ligne 15 du fichier `main.go`.

```go
var serverAddress = "localhost:8080"
```
