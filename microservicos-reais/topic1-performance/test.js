import http from "k6/http";
import { sleep } from "k6";

export const options = {
    vus: 10, // 10 usuários virtuais
    duration: "30s", // 30 segundos de teste
};

export default function () {
    const url = "http://localhost:8082/compute";
    const res = http.get(url);

    console.log(`Status: ${res.status}, Body: ${res.body}`);
    sleep(1);
}