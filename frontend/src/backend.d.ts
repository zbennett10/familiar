import { SocialUser } from "angularx-social-login";

interface Familiar {
    OnLogin(user: { Email: string, FirstName: string, LastName: string});
}

declare global {
    interface Window {
        backend: Familiar;
    }
}

