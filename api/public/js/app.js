$('#search').on('click', () => {
    ip = $('#ip').val()
    $.getJSON('http://localhost:6174/ip/' + ip, (res) => {
        $('#result').val(JSON.stringify(res, null, 4))
    })
    $('#result').val('')
    
})