import {UserModel, ProductModel} from '.';

export interface LogModel {
  id?: string;
  user?: UserModel;
  product?: ProductModel;
  action: string;
  amount?: number;
}
