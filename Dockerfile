# Spécifie l'image de base à utiliser pour l'image Docker. Dans ce cas, 
# c'est une image latest de go: derniere version.
FROM golang:1.23
# Définit le répertoire de travail à l'intérieur du conteneur s
# C'est là que les commandes suivantes seront exécutées.
# Ajoute une étiquette personnalisée à l'image
LABEL version="1.0" \
    description="Mon application forum" \
    maintainer="kouakou stephanie"
WORKDIR /app
# Copie le fichier go.mod depuis le contexte de construction (généralement le répertoire local où la commande de construction Docker est exécutée)
# dans le répertoire de travail actuel dans le conteneur.
COPY go.mod .
COPY go.sum .
# Télécharge les modules et les dépendances Go définis dans 
RUN go mod download
# Copie le reste du code source de l'application et les fichiers depuis le contexte de construction dans le répertoire de travail actuel du conteneur.
COPY . .
RUN go build -o main .
# This container exposes port 8080 to the outside world
#Informe Docker que le conteneur écoutera sur le port 8080 au moment de l'exécution. Cependant, cela ne publie pas effectivement le port, c'est plus une fonctionnalité de documentation.
EXPOSE 8000
# cmd to start service
CMD ["./main"]