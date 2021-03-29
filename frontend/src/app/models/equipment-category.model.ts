import { CommentStmt } from '@angular/compiler';
import { ItemCost } from './cost.model';


// Base interface for equipment interfaces below
interface Equipment extends EquipmentResource {
    equipment_category: string;
    weight: number;
    cost: ItemCost;
}

export interface EquipmentResource {
    index: string;
    name: string;
    url: string;
}

export interface WeaponProperty {
    index: string;
    name: string;
    desc: string[];
    url: string;
}

export interface Weapon extends Equipment {
    weapon_category: string;
    weapon_range: string;
    category_range: string;
    damage: object; // TODO
    two_handed_damage: object; // TODO
    range: object; // TODO
    properties: WeaponProperty[];
}

export interface Armor extends Equipment {
    armor_category: string;
    armor_class: object; // TODO:
    str_minimum: number;
    stealth_disadvantage: boolean;
}

export interface Gear extends Equipment {
    gear_category: string;
}

export interface EquipmentPack extends Gear {
    contents: Gear[];
}

export interface MagicItem extends Equipment {
    desc: string[]
}

export interface EquipmentCategory {
    index: string;
    name: string;
    equipment: Equipment[];
}
