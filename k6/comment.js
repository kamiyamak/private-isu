import http from "k6/http";

import { check } from 'k6';

import { parseHTML } from 'k6/html';

import { url } from './config.js';

export default function () {
    const login_res = http.post(url("/login"), {
        account_name: 'terra',
        password: 'terraterra',
    });

    check(login_res, {
        
    })
}