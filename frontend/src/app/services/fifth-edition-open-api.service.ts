import { Injectable } from "@angular/core";
import { HttpClient} from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { EquipmentCategory, EquipmentResource } from '../models/equipment-category.model';

@Injectable()
export class FifthEditionOpenAPIService {
    private baseUrl = 'https://www.dnd5eapi.co/api/'
    
    constructor(private http: HttpClient) {}

    fetchEquipmentCategories$(): Observable<EquipmentCategory[]> {
        return this.http
                   .get<{ results: EquipmentCategory[] }>(`${this.baseUrl}equipment-categories`)
                   .pipe(
                       map(({ results }) => results)
                   );
    }

    fetchAllEquipment(): Observable<EquipmentResource[]> {
        return this.http
                   .get<{ results: EquipmentResource[] }>(`${this.baseUrl}equipment`)
                   .pipe(
                       map(({ results }) => results)
                   );
    }

}
