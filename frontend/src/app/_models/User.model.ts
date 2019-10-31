export interface UserModel {
  id?: string;
  name: string;
  email?: string;
  verified: boolean;
  gravatar: boolean;
  imageUrl: string;
  balance: number;
  pin?: string;
  createdAt?: string;
  updatedAt?: string;
}
