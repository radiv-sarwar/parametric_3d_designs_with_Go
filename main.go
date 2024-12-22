package main

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/ljanyst/ghostscad/sys"

	. "github.com/ljanyst/ghostscad/primitive"
)

func main() {
	sys.Initialize()

	// ASSIGNING DIFFERENT SHAPES (IN THIS CASE PRIMITIVES) TO VARIABLES
	Cube1 := NewCube(mgl64.Vec3{5, 5, 5})
	Sphere1 := NewSphere(4)

	// PERFORMING DIFFERENT OPERATIONS WITH THE SHAPES LIKE UNION, DIFFERENCE, INTERSECTION AND HULL
	CubeUnionSphere := NewUnion(Cube1, Sphere1)
	CubeDifferenceSphere := NewDifference(Cube1, Sphere1)
	CubeIntersectSphere := NewIntersection(Cube1, Sphere1)
	CubeHullSphere := NewHull(Cube1, Sphere1)

	// TRANSFORMING THOSE SHAPES BY TRANSLATING THEM ALONG THE Y AXIS.
	// The vector values are in the format {x,y,z}
	shape2 := NewTranslation(mgl64.Vec3{0, 10, 0}, CubeDifferenceSphere)
	shape3 := NewTranslation(mgl64.Vec3{0, 20, 0}, CubeIntersectSphere)
	shape4 := NewTranslation(mgl64.Vec3{0, 30, 0}, CubeHullSphere)

	// SHOWING ALL THE OPERATIONS SIDE BY SIDE TO DEMONSTRATE DIFFERENT OPERATIONS
	Final_Shape := NewUnion(CubeUnionSphere, shape2, shape3, shape4)
	println(Final_Shape)

	// USING SLICES IN GO TO MAKE AN ARRAY OF SHAPES(DEFINED AS PRIMITIVES IN THE LIBRARY BEING USED)
	list1 := []Primitive{Cube1, shape2, shape3, shape4}

	for i := 0; i < len(list1); i++ {

		println(list1[i])
	}

	// NEW SHAPE VARIABLE ASSIGNMENT TO SHOWCASE ITERATIONS
	model1 := addCubes(Cube1)
	NewModel := spiralStair(model1)

	sys.RenderMultiple([]sys.Shape{
		{"one", Cube1, sys.None},
		{"two", NewModel, sys.Default},
		{"three", shape3, sys.None},
		{"four", shape2, sys.None},
		{"five", NewModel, sys.None},
	})

}

// ITERATION FUNCTION TO ADD MORE CUBES OF THE SAME SIZE ALONG THE Y AXIS
func addCubes(target_shape Primitive) Primitive {
	for i := 0; i < 10; i++ {
		shape := NewCube(mgl64.Vec3{5, 5, 5})
		transformed_shape := NewTranslation(mgl64.Vec3{0, (float64(i) * 5), 0}, shape)
		target_shape = NewUnion(target_shape, transformed_shape)

	}
	return target_shape

}

func stairCube(target_shape Primitive) Primitive {
	for i := 0; i < 10; i++ {
		shape := NewCube(mgl64.Vec3{5, 5, 5})
		transformed_shape := NewTranslation(mgl64.Vec3{0, (float64(i) * 5), (float64(i) * 5)}, shape)
		target_shape = NewUnion(target_shape, transformed_shape)
	}
	return target_shape
}

func spiralStair(target_shape Primitive) Primitive {
	initial_shape := target_shape
	for i := 0; i < 72; i++ {
		// transformed_shape0 := NewTranslation(mgl64.Vec3{0, (float64(i) * 2), 0}, initial_shape)
		transformed_shape1 := NewRotation(mgl64.Vec3{0, 0, 5}, initial_shape)
		transformed_shape2 := NewTranslation(mgl64.Vec3{0, 0, 5}, transformed_shape1)
		initial_shape = transformed_shape2
		target_shape = NewUnion(target_shape, transformed_shape2)
	}
	return target_shape
}
