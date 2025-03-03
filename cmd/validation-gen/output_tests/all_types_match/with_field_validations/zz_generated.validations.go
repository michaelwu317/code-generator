//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by validation-gen. DO NOT EDIT.

package withfieldvalidations

import (
	context "context"
	fmt "fmt"

	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	field "k8s.io/apimachinery/pkg/util/validation/field"
	testscheme "k8s.io/code-generator/cmd/validation-gen/testscheme"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *testscheme.Scheme) error {
	scheme.AddValidationFunc((*T1)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T1(ctx, op, nil /* fldPath */, obj.(*T1), safe.Cast[*T1](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc((*T2)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T2(ctx, op, nil /* fldPath */, obj.(*T2), safe.Cast[*T2](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	scheme.AddValidationFunc((*T4)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}, subresources ...string) field.ErrorList {
		if len(subresources) == 0 {
			return Validate_T4(ctx, op, nil /* fldPath */, obj.(*T4), safe.Cast[*T4](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresources: %v", obj, subresources))}
	})
	return nil
}

func Validate_T1(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *T1) (errs field.ErrorList) {
	// field T1.S
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *string) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field T1.S")...)
			return
		}(fldPath.Child("s"), &obj.S, safe.Field(oldObj, func(oldObj *T1) *string { return &oldObj.S }))...)

	// field T1.T2
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *T2) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field T1.T2")...)
			errs = append(errs, Validate_T2(ctx, op, fldPath, obj, oldObj)...)
			return
		}(fldPath.Child("t2"), &obj.T2, safe.Field(oldObj, func(oldObj *T1) *T2 { return &oldObj.T2 }))...)

	// field T1.T3
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *T3) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field T1.T3")...)
			return
		}(fldPath.Child("t3"), &obj.T3, safe.Field(oldObj, func(oldObj *T1) *T3 { return &oldObj.T3 }))...)

	return errs
}

func Validate_T2(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *T2) (errs field.ErrorList) {
	// field T2.S
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *string) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field T2.S")...)
			return
		}(fldPath.Child("s"), &obj.S, safe.Field(oldObj, func(oldObj *T2) *string { return &oldObj.S }))...)

	return errs
}

func Validate_T4(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *T4) (errs field.ErrorList) {
	// field T4.S
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *string) (errs field.ErrorList) {
			errs = append(errs, validate.FixedResult(ctx, op, fldPath, obj, oldObj, false, "field T4.S")...)
			return
		}(fldPath.Child("s"), &obj.S, safe.Field(oldObj, func(oldObj *T4) *string { return &oldObj.S }))...)

	return errs
}
