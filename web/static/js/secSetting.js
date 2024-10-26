const GlobNet = document.getElementById("globNet")
const Firewall = document.getElementById("Firewall")
const ReqAuth = document.getElementById("reqauth")


GlobNet.addEventListener('change', function() {
    let body = {
        "global": GlobNet.checked
    }
    console.log(body)
    sendPostRequest('/setting/api/sec/network', body, function(result) {
        const token = document.getElementById('tgToken');
        const stat = document.getElementById('system-status');
        const task = document.getElementById('tasks');
        console.log(result)
        token.value = result.token
        if (result.sysSt == true){
            stat.checked = true;
        }
        if (result.newTask == true){
            task.checked = true;
        }
    });
});

Firewall.addEventListener('change', function() {
    let body = {
        "firewall": Firewall.checked
    }
    console.log(body)
    sendPostRequest('/setting/api/sec/firewall', body, function(result) {
        const token = document.getElementById('tgToken');
        const stat = document.getElementById('system-status');
        const task = document.getElementById('tasks');
        console.log(result)
        token.value = result.token
        if (result.sysSt == true){
            stat.checked = true;
        }
        if (result.newTask == true){
            task.checked = true;
        }
    });
});

ReqAuth.addEventListener('change', function() {
    let body = {
        "auth": ReqAuth.checked
    }
    console.log(body)
    sendPostRequest('/setting/api/sec/auth', body, function(result) {
        const token = document.getElementById('tgToken');
        const stat = document.getElementById('system-status');
        const task = document.getElementById('tasks');
        console.log(result)
        token.value = result.token
        if (result.sysSt == true){
            stat.checked = true;
        }
        if (result.newTask == true){
            task.checked = true;
        }
    });
});


function SecOnlineBlock() {
    sendPostRequest('/setting/api/sec/settings', {}, function(result) {
        console.log(result)
        GlobNet.checked = result.gbn
        Firewall.checked = result.fw
        ReqAuth.checked = result.ra
        // Удаляем все классы, кроме баз
    });
}