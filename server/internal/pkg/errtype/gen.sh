#!/bin/bash

echo "// GENERATED"
echo "// DO NOT EDIT"
echo "package errtype"
echo ""

TEMPLATE=$(cat <<-END
type TYPE struct {
    Err error
}

func (e TYPE) Error() string {
    return e.Err.Error()
}
func (e TYPE) Unwrap() error {
    return e.Err
}
END
)

for var in "$@"
do
    echo -e "$TEMPLATE" | sed s/TYPE/"$var"/g
    echo ""
done