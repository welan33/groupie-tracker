//tab possède toute les coordonnées 
let tab = [];
console.log(tab)

//let res = [];

// On initialise la latitude et la longitude de Paris (centre de la carte)
let lata = 35.0;
let lone = 0.0;
let macarte = null;

// Fonction d'initialisation de la carte
function initMap() {
    // Créer l'objet "macarte" et l'insèrer dans l'élément HTML qui a l'ID "map"
    macarte = L.map('map').setView([lata, lone], 3);
    // Leaflet ne récupère pas les cartes (tiles) sur un serveur par défaut. Nous devons lui préciser où nous souhaitons les récupérer. Ici, openstreetmap.fr
    L.tileLayer('https://{s}.tile.openstreetmap.fr/osmfr/{z}/{x}/{y}.png', {
        // Il est toujours bien de laisser le lien vers la source des données
        attribution: 'données © <a href="//osm.org/copyright">OpenStreetMap</a>/ODbL - rendu <a href="//openstreetmap.fr">OSM France</a>',
        minZoom: 2,
        maxZoom: 7
    }).addTo(macarte);
    // Nous parcourons la liste des villes
    for (v in tab) {
        console.log(v)
        let marker = L.marker([tab[v].lat, tab[v].lon]).addTo(macarte);
        marker.bindPopup(tab[v].dates)
    }
}
window.onload = function () {
    // Fonction d'initialisation qui s'exécute lorsque le DOM est chargé
    initMap();
};


