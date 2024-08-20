Json = {
    JsonCheck: function (param) {
        axios.get('/api/development/check-json')
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            })
    },
    JsonFormat: function () {
        axios.get('/api/document/json-to-excel')
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            })
    },    JsonExcel: function () {
        axios.get('/api/document/json-to-excel')
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            })
    }
}


