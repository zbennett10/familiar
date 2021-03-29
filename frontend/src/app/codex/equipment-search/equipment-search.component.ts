import { Component, OnInit } from "@angular/core";
import { Observable } from 'rxjs';
import { EquipmentCategory, EquipmentResource } from 'src/app/models/equipment-category.model';
import { FifthEditionOpenAPIService } from '../../services/fifth-edition-open-api.service'

@Component({
  selector: 'familiar-equipment-search',
  templateUrl: './equipment-search.component.html',
  styleUrls: ['./equipment-search.component.scss']
})
export class EquipmentSearchComponent implements OnInit {

    equipmentResources$: Observable<EquipmentResource[]>;
    equipmentCategories$: Observable<EquipmentCategory[]>;

    selectedCategory: EquipmentCategory;
    selectedResource: EquipmentResource;

    constructor(private fifthEditionApi: FifthEditionOpenAPIService) { }

    ngOnInit() {
        this.equipmentCategories$ = this.fifthEditionApi.fetchEquipmentCategories$();
        this.equipmentResources$ = this.fifthEditionApi.fetchAllEquipment();
    }


}
