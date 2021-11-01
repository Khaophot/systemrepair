
import { EquipmentsInterface } from "./IEquipment";
import { ProblemsInterface } from "./IProblem";
import { UrgenciesInterface } from "./IUrgency";
import { CheckInsInterface } from "./ICheckIn";


export interface RepairInformationsInterface {

    ID: number,
    Datetime: Date,
    
    EquipmentID: number,
    Equipment: EquipmentsInterface,
    ProblemID: number,
    Problem: ProblemsInterface
    UrgencyID: number,
    Urgency: UrgenciesInterface
    CheckInID: number,
    CheckIn: CheckInsInterface,
   
}