import Step from './step'
type Workflow = {
  id: number;
  name: string;
  rootStepId: number;
  rootStep: Step;
  createdAt: Date;
  updatedAt: Date;
  deletedAt: Date;
};

export default Workflow;
