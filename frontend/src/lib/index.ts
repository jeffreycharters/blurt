import { browser } from "$app/environment";

export const API_ADDRESS = build_api_address()
export const WS_ADDRESS = build_ws_address()

export const BLURT_BATCH_COUNT = 50


function build_api_address() {
    if (browser) {
        const location = new URL(window.location.href)
        if (location.port) location.port = "3320"
        return location.origin + "/api"
    }
    return "http://localhost:3320/api"
}

function build_ws_address() {
    if (browser) {
        const location = new URL(window.location.href)
        if (location.port) location.port = "3320"
        location.protocol = location.protocol === "http:" ? "ws:" : "wss:"
        return location.origin + "/ws"
    }
    return "ws://localhost:3320/ws"
}


export function getDisplayDate(blurt_date: Date, current_date: Date) {
    const timeGap = +current_date - +blurt_date
    const fewSeconds = 1000 * 60;
    const fewMinutes = 1000 * 60 * 15;
    // const aboutHour = 1000 * 60 * 45;
    const oneHour = 1000 * 60 * 60;
    const hourish = 1000 * 60 * 90;
    const coupleHours = 1000 * 60 * 180;
    const oneDay = 1000 * 60 * 60 * 24;

    if (timeGap < fewSeconds) {
        return 'A few seconds ago.';
    } else if (timeGap < fewMinutes) {
        return 'A few minutes ago.';
    } else if (timeGap < oneHour) {
        return 'Less than an hour ago.';
    } else if (timeGap < hourish) {
        return 'About an hour ago.';
    } else if (timeGap < coupleHours) {
        return 'A couple of hours ago.';
    } else if (timeGap < oneDay) {
        return `About ${Math.ceil(timeGap / oneHour)} hours ago.`;
    } else {
        const options: Intl.DateTimeFormatOptions = { year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };
        return blurt_date.toLocaleDateString("en-CA", options);
    }
};