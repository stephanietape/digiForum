// Envoie du formulaire au serveur | INSCRIPTION
document.querySelector(".sign-up-form").addEventListener("submit", function (event) {
    event.preventDefault(); // Empêche l'envoi du formulaire par défaut

    // Récupérer les valeurs des champs du formulaire
    var username = document.querySelector(".username").value;
    var email = document.querySelector(".email").value;
    var password = document.querySelector(".password").value;

    // Créer un objet contenant les données du formulaire
    var formData = {
        username: username,
        email: email,
        password: password,
        nom: "",
        prenom: "",
        photo: "lol.png"
    };

    console.log(formData)

    // Envoyer les données du formulaire au serveur Go via une requête POST
    fetch("/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => {
            if (!response.ok) {
                var notificationDiv = document.getElementById("notification");
                notificationDiv.textContent = "Cet e-mail ou nom d'utilisateur est déjà pris utilisé";
                notificationDiv.style.backgroundColor = "#ED3434"
                notificationDiv.style.opacity = 0.80;
                setTimeout(function () {
                    notificationDiv.textContent = ""; // Effacer le contenu de la div
                    notificationDiv.style.opacity = 0;
                }, 5000);
                throw new Error("Erreur lors de l'envoi des données");

            }
            // Ne pas essayer de parser la réponse comme JSON, car elle n'est pas JSON
            return response.text(); // Récupérer le texte de la réponse
        })
        .then(data => {
            // Afficher la réponse du serveur
            console.log("Réponse du serveur:", data);
            // Afficher une notification de succès

            // Afficher une notification de succès
            var notificationDiv = document.getElementById("notification");
            notificationDiv.textContent = "Inscription réussie ! vous pouvez vous connecter ";
            notificationDiv.style.backgroundColor = "#48A84C"
            notificationDiv.style.opacity = 0.80;
            //vider les champs
            document.querySelector(".username").value = "";
            document.querySelector(".email").value = "";
            document.querySelector(".password").value = "";
            // Effacer la notification après 3 secondes
            setTimeout(function () {
                notificationDiv.textContent = ""; // Effacer le contenu de la div
                notificationDiv.style.opacity = 0;
            }, 5000); // Durée en millisecondes (3 secondes dans cet exemple)


        })
        .catch(error => {
            console.error("Erreur:", error.message);


        });
});

// Envoie du formulaire au serveur | CONNEXION
document.querySelector(".sign-in-form").addEventListener("submit", function (event) {
    event.preventDefault(); // Empêche l'envoi du formulaire par défaut

    // Récupérer les valeurs des champs du formulaire
    var username = document.querySelector(".usernamePass").value;
    var password = document.querySelector(".motdepasse").value;
    // Créer un objet contenant les données du formulaire
    var formData = {
        username: username,
        password: password,
    };

    console.log(formData)
    fetch("/connexion", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => {
            if (!response.ok) {
                var notificationDiv = document.getElementById("notificationCon");
                notificationDiv.textContent = "Mot de passe ou username incorrect";
                notificationDiv.style.backgroundColor = "#ED3434"
                notificationDiv.style.opacity = 0.80;
                setTimeout(function () {
                    notificationDiv.textContent = ""; // Effacer le contenu de la div
                    notificationDiv.style.opacity = 0;
                }, 5000);
                throw new Error("Erreur lors de l'envoi des données");

            }
            // Ne pas essayer de parser la réponse comme JSON, car elle n'est pas JSON
            return response.text(); // Récupérer le texte de la réponse
        })
        .then(data => {
            // Afficher la réponse du serveur
            console.log("Réponse du serveur:", data);

            // Afficher une notification de succès
            var notificationDiv = document.getElementById("notificationCon");
            notificationDiv.textContent = "connexion réussie !  ";
            notificationDiv.style.backgroundColor = "#48A84C"
            notificationDiv.style.opacity = 0.80;
            //vider les champs
            document.querySelector(".usernamePass").value = "";
            document.querySelector(".motdepasse").value = "";
            // Effacer la notification après 3 secondes
            setTimeout(function () {
                notificationDiv.textContent = ""; // Effacer le contenu de la div
                notificationDiv.style.opacity = 0;
            }, 1000); // Durée en millisecondes (3 secondes dans cet exemple)


        })
        .catch(error => {
            console.error("Erreur:", error.message);


        });
});
