import http from "k6/http";
import { check, sleep } from "k6";

export let options = {
  stages: [
    { duration: "10s", target: 10 },
    { duration: "20s", target: 50 },
    { duration: "30s", target: 100 },
    { duration: "30s", target: 0 },
  ],
  thresholds: {
    http_req_failed: ["rate==0"],
    http_req_duration: ["p(95)<200"],
  },
};

export default function () {
  let res = http.get("http://localhost:3000/api/categories", {
    headers: {
      Accept: "application/json",
      "X-API-Key": "RAHASIA",
    },
  });

  check(res, {
    "status 200": (r) => r.status === 200,
    "durasi < 200ms": (r) => r.timings.duration < 200,
  });
}
