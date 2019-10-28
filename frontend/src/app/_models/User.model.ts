export interface UserModel {
  id?: string;
  name: string;
  email?: string;
  verified: boolean;
  gravatar: boolean;
  imageUrl: boolean;
  balance: number;
  pin?: string;
  createdAt?: string;
  updatedAt?: string;
}
