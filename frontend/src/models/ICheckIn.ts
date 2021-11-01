import { CustomersInterface } from "./ICustomer";
import { RoomsInterface } from "./IRoom";

export interface CheckInsInterface {

    ID: number,

    RoomID: number,
    Room: RoomsInterface,
    
    CustomerID: number,
    Customer: CustomersInterface,

   }