package Domain

type UserI interface {
	getAll()
	getOne()
	create()
	update()
	delete()
}
