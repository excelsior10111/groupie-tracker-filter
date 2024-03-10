var selCountries = [];  
var element = document.getElementsByClassName("card");
const rangeInput = document.querySelectorAll(".range-input input");
yearInput = document.querySelectorAll(".year-input input");
progress = document.querySelector(".slider .progress");
let priceGap = 2
let intArray = [];
var checkboxes = document.querySelectorAll("input[type='checkbox']");
let checked = false

function checkAll(mycheckbox) {
    if (mycheckbox.checked == true) {
        checkboxes.forEach(function(checkbox) {
            checkbox.checked = true;
        });
    } else {
        checkboxes.forEach(function(checkbox) {
            checkbox.checked = false;
        });
    }
}
yearInput.forEach(input => {
    input.addEventListener("input", e => {
        //getting two inputs and parsing them to number
        let minVal = parseInt(yearInput[0].value),
            maxVal = parseInt(yearInput[1].value);
        intArray = [];
        for (let i = 0; i <= 51; i++) {
            intArray.push(i);
        }
        for (var i = 0; i < element.length; i++) {
            if (artists[i].creationDate >= minVal && artists[i].creationDate <= maxVal) {
                element[i].style.display = "block";
                for (var i = 0; i < checkboxes.length; i++) {
                    checkboxes[i].addEventListener('change', function() {});
                }
            } else {
                intArray.push(i);
                element[i].style.display = "none"; // Hide the element
            }
        }
        if ((maxVal - minVal >= priceGap) && maxVal <= 2023) {
            if (e.target.className === "input-min") { //if active input is min input
                rangeInput[0].value = minVal;
                progress.style.left = (minVal / rangeInput[0].max) * 100 + "%";
            } else {
                rangeInput[1].value = maxVal
                progress.style.right = 100 - (maxVal / rangeInput[1].max) * 100 + "%";
            }
        }
        checkboxes.forEach(function(checkbox) {
            checkbox.checked = false;
        });
    });
});
rangeInput.forEach(input => {
    input.addEventListener("input", e => {
        checked = true
        let minVal = parseInt(rangeInput[0].value)
        maxVal = parseInt(rangeInput[1].value);
        intArray = [];
        if (maxVal - minVal < priceGap) {
            if (e.target.className === "range-min") {
                rangeInput[0].value = maxVal - priceGap;
            } else {
                rangeInput[1].value = minVal + priceGap
            }
        } else {
            yearInput[0].value = minVal;
            yearInput[1].value = maxVal
            progress.style.left = (minVal / rangeInput[0].max) * 100 + "%";
            progress.style.right = 100 - (maxVal / rangeInput[1].max) * 100 + "%";
        }
        for (var i = 0; i < element.length; i++) {
            if (artists[i].creationDate >= minVal && artists[i].creationDate <= maxVal) {
                intArray.push(i);
                element[i].style.display = "block"; 
            } else {
                element[i].style.display = "none"; 
            }
        }
        myMap.forEach((value, key) => {
            myMap.set(key, []);
          });
        handleCheckboxChanges()
        console.log("Called from range func")
    })
})
if (!checked){
    checkboxes.forEach(function(checkbox) {
        checkbox.checked = true;
    });
}
let myMap = new Map();
for (let i = 1; i <= 8; i++) {
    myMap.set(i, []);
}
for (var i = 0; i < checkboxes.length; i++) {
    checkboxes[i].addEventListener('change', function() {
        console.log("Called from select func")
        handleCheckboxChanges()
    });
}
if (selCountries.length > 0){
    console.log(selCountries)
    for (var i = 0; i < checkboxes.length; i++) {
        checkboxes[i].addEventListener('change', function() {
            console.log("Called from select func")
            handleCheckboxChanges()
        });
    }
}
new MultiSelectTag('countries', { 
    onChange: function(values) {
        selCountries = [];
        checked=false
        for (j=0; j<values.length; j++){
            selCountries.push(values[j].value)
        }
        handleCheckboxChanges()
}}) 
function handleCheckboxChanges() {
    console.log("new range")
    if (!checked){
        for (let i = 0; i <= 51; i++) {
            intArray.push(i);
        }
        checked = true
    }
    console.log(intArray)
    for (var j = 0; j < checkboxes.length; j++) {
        if (checkboxes[j].checked) {
            for (var k = 0; k < intArray.length; k++) {
                var memberLen = artists[intArray[k]].members.length;
                mapCheck = new Map();
                for (d=0; d < selCountries.length; d++){
                    mapCheck.set(selCountries[d], false)
                }
                num = selCountries.length
                if (memberLen == j) {
                    if (selCountries.length > 0) {
                        console.log(selCountries.length)
                        for (m=0; m < artists[intArray[k]].locations2.length; m++){
                            var country = artists[intArray[k]].locations2[m].indexOf("-")
                            let tempresult = artists[intArray[k]].locations2[m].slice(country+1, artists[intArray[k]].locations2[m].length).replace(/_/g, ' ');
                            var words = tempresult.split(' ');
                            for (var n = 0; n < words.length; n++) {
                                words[n] = words[n][0].toUpperCase() + words[n].substr(1);
                            }
                            var result =  words.join(' ');
                            if (mapCheck.has(result)){
                                mapCheck.set(result, true)
                            }
                        }
                        mapCheck.forEach(value => {
                            if (value == true){
                                num--
                            }
                        });
                        if (num == 0){
                            console.log(artists[intArray[k]].name, num)
                            myMap.get(j).push(intArray[k]);
                            element[intArray[k]].style.display = "block"; // Show the element
                            intArray.splice(k, 1)
                            k--;
                        }

                    }else{
                        myMap.get(j).push(intArray[k]);
                        element[intArray[k]].style.display = "block"; // Show the element
                        intArray.splice(k, 1)
                        k--;
                    }
                } else {
                    element[intArray[k]].style.display = "none"; // Hide the element
                }
            }
        } else {
            if (myMap.has(j)) {
                let mapIntArray = myMap.get(j);
                intArray = [...intArray, ...mapIntArray];
                myMap.set(j, []);
                console.log(myMap.get(j));
                console.log(intArray);

            }
            for (var k = 0; k < artists.length; k++) {
                var memberLen = artists[k].members.length;
                if (memberLen == j) {
                    element[k].style.display = "none"; // Hide the element
                }
            }
        }
    }
}