//getting all required elements
const searchWrapper = document.querySelector(".search-input");
const inputBox = searchWrapper.querySelector("input");
const suggBox = document.querySelector(".autocom-box");

//if user press any key and release
inputBox.onkeyup = (e)=>{
    let userData = e.target.value; //user entered data
    let emptyArray = [];
    if(userData){
        emptyArray = suggestions.filter((data)=>{
            //filtering array value and user char to lowercase and return only those word/sentence 
            //which are starts with user entered word
            return data.toLocaleLowerCase().startsWith(userData.toLocaleLowerCase());
        });
        emptyArray = emptyArray.map((data)=>{
            //return data = '<li>' + data + '</li>';
            return data = '<form action="/artist" method="post" autocomplete="off">' + '<button type="submit" class="btn" style="display: relative;">' + data + '</button>' + '<input id="prodId" name="artist" type="hidden" value={{ .Id }}>' + '</form>';
        });
        console.log(emptyArray);
        searchWrapper.classList.add("active"); //show autocomplete box
        showSuggestions(emptyArray);
        let allList = suggBox.querySelectorAll("button");
        for (let i = 0; i < allList.length; i++) {
            //adding onclick attribute in all li tag
            allList[i].setAttribute("onclick", "select(this)");
            
        }
    }else{
        searchWrapper.classList.remove("active"); //hide autocomplete box
    }
}

//associer id à la suggestion

function select(element) {
    let selectUserData = element.textContent;
    inputBox.value = selectUserData; //passing the user selected list item data in textfield (placeholder)
    searchWrapper.classList.remove("active");
}

function showSuggestions(list) {
    let listData;
    if(!list.length){
        userValue = inputBox.value;
        //listData = '<li>' + userValue + '</li>';
        listData = '<form action="/artist" method="post" autocomplete="off">' + '<button type="submit" class="btn" style="display: relative;">' + userValue + '</button>' + '<input id="prodId" name="artist" type="hidden" value={{ .Id }}>' + '</form>'
    }else{
        listData = list.join('');
    }
    suggBox.innerHTML = listData;
}