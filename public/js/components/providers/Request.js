class Request{
    constructor(){}

    Get(url) {
        return fetch(url, {
            method: 'GET'
        }).then(res => res.json())
    }

    Post(url, data) {
        return fetch(url, {
            method: 'POST',
            body: JSON.stringify(data),
            headers:{
                'Content-Type': 'application/json'
            }
        }).then(res => res.json())
    }

    Put(url, data) {
        return fetch(url, {
            method: 'PUT',
            body: JSON.stringify(data),
            headers:{
                'Content-Type': 'application/json'
            }
        }).then(res => res.json())
    }

    Delete(url) {
        return fetch(url, {
            method: 'DELETE'
        }).then(res => res.json())
    }
}

export default Request;