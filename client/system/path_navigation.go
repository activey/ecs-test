package system

import (
	"ecs-test/assets/sprites/animations"
	"ecs-test/client/component"
	"ecs-test/client/component/world"
	"fmt"
	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
	"image/color"
	"log"
)

type PathNavigation struct {
	worldMapQuery *donburi.Query
	debugQuery    *donburi.Query

	playerEntry   *donburi.Entry
	playerLock    *donburi.Entry
	worldMapEntry *donburi.Entry
	cameraEntry   *donburi.Entry

	startNode   *world.NavigationNode
	endNode     *world.NavigationNode
	currentPath *world.NavigationPath

	debugEntry    *donburi.Entry
	flagAnimation *ganim8.Animation
}

func NewPathNavigation() *PathNavigation {
	return &PathNavigation{
		worldMapQuery: donburi.NewQuery(filter.Contains(component.WorldMap)),
		debugQuery:    donburi.NewQuery(filter.Contains(component.Debug)),

		flagAnimation: animations.NewFlagAnimation(),
	}
}

func (n *PathNavigation) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	if n.cameraEntry == nil || n.worldMapEntry == nil {
		return
	}
	worldMapData := component.WorldMap.Get(n.worldMapEntry)
	if !worldMapData.IsLoaded() {
		return
	}

	cameraTransform := transform.Transform.Get(n.cameraEntry)
	if n.currentPath != nil && n.currentPath.HasNodes() {
		walkableLayer := worldMapData.WalkableLayer

		if n.endNode != nil {
			walkableLayer.DrawNavigationPath(n.currentPath, color.RGBA{R: 255, G: 215, B: 0, A: 255}, screen, cameraTransform)
			walkableLayer.DrawTargetFlag(n.endNode.Vec2(), n.flagAnimation, screen, cameraTransform)
		}
	}

	worldMap := component.WorldMap.Get(n.worldMapEntry)
	debug := component.Debug.Get(n.debugEntry)
	if !debug.IsEnabled() {
		return
	}
	if n.endNode != nil {
		n.endNode.DrawDebug(
			screen,
			cameraTransform,
			float64(worldMap.WalkableLayer.TileWidth),
			float64(worldMap.WalkableLayer.TileHeight),
		)
	}
}

func (n *PathNavigation) Update(ecs *ecs.ECS) {
	n.findDebugComponent(ecs)
	n.findPlayer(ecs)
	n.findPlayerLock(ecs)
	n.findWorldMap(ecs)
	n.findCamera(ecs)

	if n.playerEntry == nil || n.playerLocked() {
		return
	}

	n.handleClick(ecs)
	n.updateFlagIfNecessary()
}

func (n *PathNavigation) playerLocked() bool {
	if n.playerLock == nil {
		return true
	}

	lock := component.PlayerLock.Get(n.playerLock)
	return lock.IsLocked()
}

func (n *PathNavigation) handleClick(ecs *ecs.ECS) {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mouseX, mouseY := ebiten.CursorPosition()
		playerX, playerY := n.playerPosition()

		cameraTransform := transform.Transform.Get(n.cameraEntry)
		worldMap := component.WorldMap.Get(n.worldMapEntry)

		n.startNode = worldMap.WalkableLayer.GetNodeFromPixel(
			playerX,
			playerY,
			cameraTransform,
		)
		n.endNode = worldMap.WalkableLayer.GetNodeFromPixel(
			float64(mouseX),
			float64(mouseY),
			cameraTransform,
		)

		if n.endNode != nil && n.startNode != nil {
			path, distance, found := astar.Path(n.startNode, n.endNode)

			if found {
				fmt.Printf("Path found!, distance: %f\n", distance)
				navigationPath := world.NewNavigationPath()
				for _, node := range path {
					navigationNode := node.(*world.NavigationNode)
					navigationPath.AddNode(navigationNode)
				}
				navigationPath.ComputeOrientations()

				component.Player.Get(n.playerEntry).FollowPath(navigationPath)
				component.Camera.Get(n.cameraEntry).FollowPlayer()
				n.currentPath = navigationPath
			} else {
				n.currentPath = nil
				println("No path found")
			}
		}
	}
}

func (n *PathNavigation) playerPosition() (float64, float64) {
	cameraTransform := transform.Transform.Get(n.cameraEntry)
	playerTransform := transform.Transform.Get(n.playerEntry)

	return playerTransform.LocalPosition.
		Sub(cameraTransform.LocalPosition).
		Mul(cameraTransform.LocalScale).
		XY()
}

func (n *PathNavigation) findPlayer(ecs *ecs.ECS) {
	if n.playerEntry == nil {
		entry, found := component.PlayerQuery.First(ecs.World)
		if found {
			n.playerEntry = entry
		} else {
			// report it
		}
	}
}

func (n *PathNavigation) findPlayerLock(e *ecs.ECS) {
	if n.playerLock == nil {
		lockEntry, found := component.PlayerLockQuery.First(e.World)
		if !found {

		}
		n.playerLock = lockEntry
	}
}

func (n *PathNavigation) findWorldMap(e *ecs.ECS) {
	if n.worldMapEntry == nil {
		worldMapEntry, worldMapFound := n.worldMapQuery.First(e.World)
		if !worldMapFound {
			// Handle case where the world map entity is not found
		}
		n.worldMapEntry = worldMapEntry
	}
}

func (n *PathNavigation) findCamera(e *ecs.ECS) {
	if n.cameraEntry == nil {
		cameraEntry, cameraFound := component.CameraQuery.First(e.World)
		if !cameraFound {
			// Handle case where the world map entity is not found
		}
		n.cameraEntry = cameraEntry
	}
}

func (n *PathNavigation) findDebugComponent(e *ecs.ECS) {
	if n.debugEntry == nil {
		debugEntry, entryFound := n.debugQuery.First(e.World)
		if !entryFound {
			log.Fatalf("Debug entry not found!")
			return
		}
		n.debugEntry = debugEntry
	}
}

func (n *PathNavigation) updateFlagIfNecessary() {
	if n.currentPath != nil && n.currentPath.HasNodes() {
		n.flagAnimation.Update()
	}
}
