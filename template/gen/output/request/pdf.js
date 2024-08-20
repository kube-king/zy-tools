Pdf = {
    WordToPdf: function (param) {
        console.log(param)
        axios.get('/api/development/check-json')
            .then(function (response) {
                console.log(response);
            })
            .catch(function (error) {
                console.log(error);
            })
    },

}



