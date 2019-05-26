package mvvm

import (
	"./utils"
)

type ViewControllerCreator struct{}

func (r ViewControllerCreator) CreateDotHFile() {
	content := `//
//  __prefix__ViewController.h
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "IGCBaseViewController.h"

NS_ASSUME_NONNULL_BEGIN

@interface __prefix__ViewController : IGCBaseViewController

@end

NS_ASSUME_NONNULL_END
`
	utils.WriteFile(utils.FileTypeOfViewController, true, content)

}

func (r ViewControllerCreator) CreateDotMFile() {
	content := `//
//  __prefix__ViewController.m
//  __proj__
//
//  Created by __user__ on __time__.
//  Copyright © __year__ __user__. All rights reserved.
//

#import "__prefix__ViewController.h"
#import "__prefix__Pipeline.h"
#import "__prefix__ViewModel.h"
#import "__prefix__View.h"

#define COMPONENT_NAME @"__component__name__"

@interface __prefix__ViewController ()

@property (nonatomic, strong) __prefix__Pipeline *pipeline;

@end

@implementation __prefix__ViewController

@dynamic pipeline;

- (instancetype)initWithContext:(NSDictionary *)context {
    self = [super initWithContext:context];
    if (self != nil) {

    }
    return self;
}

- (void)viewDidLoad {
    [super viewDidLoad];
    self.view.backgroundColor = [UIColor whiteColor];
}

- (void)addObservers {
    
}

#pragma mark - IGCComponentRegisterProtocol

- (NSString *)componentName {
    return COMPONENT_NAME;
}

+ (NSDictionary<NSString *, IGCComponentConfig *> *)createComNameKeyAndComConfigObj {
    return @{COMPONENT_NAME: [IGCComponentConfig quickInitComponentWithViewModelClassName:NSStringFromClass([__prefix__ViewModel class]) viewClassName:NSStringFromClass([__prefix__View class])]};
}

@end
`
	utils.WriteFile(utils.FileTypeOfViewController, false, content)
}
