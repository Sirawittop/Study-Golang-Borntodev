const baseURL = 'https://f83e-202-80-249-211.ap.ngrok.io/course';
fetch(baseURL)
    .then(resp => resp.json())
    .then(data => displayData(data));


function displayData(data) {
    document.querySelector("pre").innerHTML = JSON.stringify(data,null ,2)
}