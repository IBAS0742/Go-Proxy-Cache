import {Message} from "element-ui";

const host = "http://localhost:8090"

const format = r => {
    if (!r.Data) {
        return r;
    }
    let rd = r.Data;
    try {
        r.Data = JSON.parse(r.Data);
    } catch (e) {
        console.log(e);
        r.Data = rd;
    }
    return r;
}
const urlApi = `${host}/api`;
export const _api = (methodName,SourceParams = [],params = {},method='post',url=urlApi) => {
    return new Promise(s => {
        return fetch(url,{
            method: method,
            body: JSON.stringify({
                MethodName: methodName,
                SourceParams: (SourceParams || []).map(_ => _ + ''),
                Params: params,
            })
        }).then(_=>_.json()).then(format).then(_ => {
            if (_.Code === 200) {
                s(_);
            } else {
                Message.error(_.Error)
                // return Promise.reject(_.Error);
            }
        });
    });
};
// export const _proxyApi = ()
const _proxyApiUrl = `${host}/ProxyApi`
export const getProxyApiTypes = {
    all: 'all',
    new: 'new',
    old: 'old'
}
export const getProxyApi = (getProxyApiType,parentId) => {
    return _api('post',[getProxyApiType,parentId],{},'post',_proxyApiUrl);
}
