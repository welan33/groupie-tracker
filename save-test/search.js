let res = [];

function search() {
    let searchinput = document.getElementById("searchInput");
    let sug = document.getElementById("suggestions")

    sug.innerHTML = ""

    const input = searchinput.value;

    for (let i = 0; i < res.length; i++) {
    if (res[i].toLowerCase().includes(input)) {
            sug.innerHTML += res[i] ;
        }
    }
    
}