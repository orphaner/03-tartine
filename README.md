# tartine, build mono repo

Tout est dans le même dépôt. 

Il y a un module pour le package core

Il y a un module par microservice. Le replace en relatif est fonctionnel.

## Démo 3
```bash
# Définition d'un GOPATH clean pour voir ce qu'il s'y passe
$ export GOPATH=/home/nicolas/go-03-tartine
$ mkdir $GOPATH
$ sudo rm -rf $GOPATH

# tartine-name marche bien en CLI
$ cd cmd/tartine-name
$ go run main.go
$ tree $GOPATH
```
ouvrir l'IDE, browser le code (go.mod en particulier),
éxecuter tartine-name, updater logger.go, vérifier que ça fonctionne :)
 

éxécuter tartine-giphy/main.go, vérifier l'erreur dans le go.mod

## Démo 4
changer la dépendance gin-gonic de 1.6.3 vers 1.6.2. Montrer que ça update tout le temps la lib, que en fait pour
figer la version il faut à tout prix passer par le replace ...
