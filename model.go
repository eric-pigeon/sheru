package sheru

type Model struct {
	ID uint `sheru:"primary_key"`
}

func Change(model interface{}) Changeset {
	return Changeset{
		metadata: schemaDataFor(model),
	}
}
