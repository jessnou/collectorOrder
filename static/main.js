document.getElementById("form").addEventListener("submit", function(event) {
    event.preventDefault();

    // Получить данные из формы
    let orderID = document.getElementById("orderID").value;
    let numbersArray = orderID.split(',').map(Number);
    fetch(`/get?ids=${numbersArray.join(',')}`)
        .then(response => response.json())
        .then(data => {
            alert(`${JSON.stringify(data)}`);
        })
        .catch(error => {

            console.error("Ошибка при запросе данных на сервер: " + error);
        });

});