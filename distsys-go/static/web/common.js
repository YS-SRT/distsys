
const baseUrl = () => {

    return "http://localhost:8090/api/v1/";
}

async function apiGet(path, onSuccess, onFailure) {
    try {
        const url = new URL(path, baseUrl());
        const resp = await fetch(url, {
            method: "GET",
            mode: "cors",
            cache: "no-cache",
            redirect: "follow",
            credentials: "omit",
            keepalive: true,
            headers: {
                "Content-Type": "application/json",
            },
        });

        if (!resp.ok) {

            onFailure(new Error('${resp.status}:${resp.statusText}'))
        }
        const data = await resp.json();
        onSuccess(data);

    } catch (err) {

        onFailure(err)
    }
}


async function apiPost(path, inputData, onSuccess, onFailure) {
    try {
        const url = new URL(path, baseUrl());
        const resp = await fetch(url, {
            method: "POST",
            mode: "cors",
            cache: "no-cache",
            redirect: "follow",
            credentials: "omit",
            keepalive: true,
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(inputData),
        });

        if (!resp.ok) {

            onFailure(new Error("${resp.status}:${resp.statusText}"))
        }

        const data = await resp.json();
        onSuccess(data);

    } catch (err) {

        onFailure(err)
    }
}

async function apiPost2(path, inputData, onSuccess, onFailure) {

    const url = new URL(path, baseUrl());
    fetch(url, {
        method: "POST",
        mode: "cors",
        cache: "no-cache",
        redirect: "follow",
        credentials: "omit",
        keepalive: true,
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(inputData),
    })
        .then(resp => {

            if (!resp.ok) {

                onFailure(new Error("${resp.status}:${resp.statusText}"))
            } 
            return resp.json()

        })
        .then(data => onSuccess(data))
        .catch(error => onFailure(error))

}








