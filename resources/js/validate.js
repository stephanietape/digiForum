// Envoie du formulaire au serveur
document.querySelector(".sign-up-form").addEventListener("submit", function(event) {
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
        prenom:"",
        photo:"lol.png"
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
            throw new Error("Erreur lors de l'envoi des données");
        }
        // Ne pas essayer de parser la réponse comme JSON, car elle n'est pas JSON
        return response.text(); // Récupérer le texte de la réponse
    })
    .then(data => {
        // Afficher la réponse du serveur
        console.log("Réponse du serveur:", data);
    })
    .catch(error => {
        console.error("Erreur:", error);
    });
});
