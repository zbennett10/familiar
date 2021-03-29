import { Component, OnInit } from "@angular/core";
import { AuthService, FacebookLoginProvider, SocialUser } from "angularx-social-login";

@Component({
  selector: 'familiar-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

    user: SocialUser;
    isLoggedIn = false;

    constructor(private authService: AuthService) { }

    ngOnInit() {
        this.authService.authState.subscribe((user) => {
            this.user = user;
            this.isLoggedIn = user ? true : false;
            if (this.isLoggedIn) {
                console.log('Notifying backend of successful login...');
                // TODO: getting error on the backend...
                window.backend.OnLogin({ Email: user.email, FirstName: user.firstName, LastName: user.lastName });
            }
        });
    }

    signInWithFB(): void {
        this.authService.signIn(FacebookLoginProvider.PROVIDER_ID);
    } 

    signOut(): void {
        this.authService.signOut();
    }

}