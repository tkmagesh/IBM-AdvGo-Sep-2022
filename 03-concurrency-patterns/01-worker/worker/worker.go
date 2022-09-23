package worker

type Work interface {
	Task()
}

type Worker struct {
	/*  */
}

func New( /*  */ ) Worker {
	return Worker{}
}

func (w *Worker) Run(work Work) {
	/*  */
}

func (w *Worker) Shutdown() {

}
