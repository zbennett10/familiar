import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgSelectModule } from '@ng-select/ng-select';
import { FormsModule } from '@angular/forms';



import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { APP_BASE_HREF } from '@angular/common';
import { SocialLoginModule, AuthServiceConfig } from "angularx-social-login";
import { FacebookLoginProvider } from "angularx-social-login";
import { LoginComponent } from './user/login.component';
import { HttpClientModule } from '@angular/common/http';
import { EquipmentSearchComponent } from './codex/equipment-search/equipment-search.component';
import { FifthEditionOpenAPIService } from './services/fifth-edition-open-api.service';

const config = new AuthServiceConfig([
  {
    id: FacebookLoginProvider.PROVIDER_ID,
    provider: new FacebookLoginProvider("1649400632114258")
  }
]);

export function provideConfig() {
  return config;
}

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    EquipmentSearchComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    SocialLoginModule,
    NgSelectModule,
    FormsModule
  ],
  providers: [
    {provide: APP_BASE_HREF, useValue : '/' },
    { provide: AuthServiceConfig, useFactory: provideConfig },
    FifthEditionOpenAPIService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
