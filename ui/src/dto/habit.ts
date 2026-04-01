export type HabitDto = {
  uuid: string;
  icon: number;
  name: string;
};

export type CreateHabitDto = {
  icon: number;
  name: string;
};

export type UpdateHabitDto = {
  icon: number;
  name: string;
};

export type ValidationErrorsDto = Record<string, string>;
